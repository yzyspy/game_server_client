package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"zinx-lib/ziface"
	"zinx-lib/zpack"
)

type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	ExitChann chan bool

	// Buffered channel used for message communication between the read and write goroutines
	// (有缓冲管道，用于读、写两个goroutine之间的消息通信)
	MsgBuffChan chan []byte

	//router管理器
	msgHandler ziface.IMsgHandle

	// Data packet packaging method
	// (数据报文封包方式)
	packet ziface.IDataPack

	// Which Connection Manager the current connection belongs to
	// (当前连接是属于哪个Connection Manager的)
	connManager ziface.IConnManager

	// Connection properties
	// (连接属性)
	property map[string]interface{}

	// Lock to protect the current property
	// (保护当前property的锁)
	propertyLock sync.Mutex

	// Hook function when the current connection is created
	// (当前连接创建时Hook函数)
	onConnStart func(conn ziface.IConnection)

	// Hook function when the current connection is disconnected
	// (当前连接断开时的Hook函数)
	onConnStop func(conn ziface.IConnection)
}

func (c *Connection) Start() {
	go c.StartRead()
	go c.StartWrite()
}

func (c *Connection) StartRead() {
	dp := zpack.DataPack{}

	for {
		binaryHead := make([]byte, dp.GetHeadLen())
		//前面8个字节是消息头,(int32 DateLen + int32 MsgID , 一共8个字节)
		if _, err := io.ReadFull(c.Conn, binaryHead); err != nil {
			fmt.Println("read head err:", err)
			continue
		}

		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("unpack head err1:", err)
			continue
		}

		fmt.Printf("recv msg , msgId = %d, dataLen = %d \n", msgHead.GetMsgID(), msgHead.GetDataLen())

		msg := msgHead.(*zpack.Message)

		//再根据DateLen进行第二次读取, 将data取出来
		msg.Data = make([]byte, msgHead.GetDataLen())

		if _, err := io.ReadFull(c.Conn, msg.Data); err != nil {
			fmt.Println("read data err:", err)
			continue
		}
		req := Request{
			conn: c,
			msg:  msg,
		}
		//go c.msgHandler.DoMsgHandler(&req, 1)
		c.msgHandler.SendMsgToTaskQueue(&req)
	}
}

func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}
	c.Conn.Close()
	close(c.ExitChann)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnID
}

// 如果返回值是接口，那么不需要返回指针了，接口本身就是指针，不要返回 *net.Addr了
// 返回接口值（版本1）是标准做法
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) StartWrite() {
	for {
		select {
		case msg := <-c.MsgBuffChan:
			err := c.Send(msg)
			if err != nil {
				fmt.Printf("SendMsg err msg ID = %d, data = %+v, err = %+v", msg, string(msg), err)
			}
		case <-c.ExitChann:
		}
	}
}

func newServerConn(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:        conn,
		ConnID:      connID,
		msgHandler:  msgHandler,
		isClosed:    false,
		ExitChann:   make(chan bool, 1),
		MsgBuffChan: make(chan []byte),
		packet:      zpack.NewDataPack(),
	}
	c.connManager = server.GetConnMgr()

	c.connManager.Add(c)

	// Hook connect

	return c
}

func (c *Connection) Send(data []byte) error {
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}
	_, err := c.Conn.Write(data)
	if err != nil {
		fmt.Printf("SendMsg err data = %+v, err = %+v", data, err)
		return err
	}
	return nil
}

// SendMsg directly sends Message data to the remote TCP client.
// (直接将Message数据发送数据给远程的TCP客户端)
func (c *Connection) SendMsg(msgID uint32, data []byte) error {

	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}
	// Pack data and send it
	msg, err := c.packet.Pack(zpack.NewMsgPackage(msgID, data))
	if err != nil {
		fmt.Printf("Pack error msg ID = %d", msgID)
		return errors.New("Pack error msg ")
	}
	//读写分离，给客户端写消息使用单独的协程
	c.MsgBuffChan <- msg
	//c.Send(msg)
	fmt.Printf("send msg , msgId = %d, dataLen = %d \n", msgID, len(data))
	return nil
}
