package znet

import (
	"fmt"
	"io"
	"net"
	"zinx-lib/ziface"
	"zinx-lib/zpack"
)

type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	ExitChann chan bool

	Router ziface.IRouter
}

func (c *Connection) Start() {
	go c.StartRead()
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

		c.Router.PreHandle(&req)
		c.Router.Handle(&req)
		c.Router.PostHandle(&req)
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

func (c *Connection) Send(data []byte) error {
	//TODO implement me
	panic("implement me")
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		Router:    router,
		isClosed:  false,
		ExitChann: make(chan bool, 1),
	}
}

//
//// SendMsg directly sends Message data to the remote TCP client.
//// (直接将Message数据发送数据给远程的TCP客户端)
//func (c *Connection) SendMsg(msgID uint32, data []byte) error {
//
//	if c.isClosed() == true {
//		return errors.New("connection closed when send msg")
//	}
//	// Pack data and send it
//	msg, err := c.packet.Pack(zpack.NewMsgPackage(msgID, data))
//	if err != nil {
//		zlog.Ins().ErrorF("Pack error msg ID = %d", msgID)
//		return errors.New("Pack error msg ")
//	}
//
//	err = c.Send(msg)
//	if err != nil {
//		zlog.Ins().ErrorF("SendMsg err msg ID = %d, data = %+v, err = %+v", msgID, string(msg), err)
//		return err
//	}
//
//	return nil
//}
