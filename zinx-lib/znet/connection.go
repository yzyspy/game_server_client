package znet

import (
	"fmt"
	"io"
	"net"
	"zinx-lib/ziface"
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
	buf := make([]byte, 1024)
	for {
		_, errRead := c.Conn.Read(buf)
		if errRead != nil {
			if errRead == io.EOF {
				//	fmt.Println("Connection closed by remote peer.")
				continue // 跳出循环，优雅地处理连接关闭
			}
			fmt.Println("Read err:", errRead)
			continue
		}
		//	c.FuncApi(c.Conn, buf[:cnt], cnt)
		req := Request{
			conn: c,
			data: buf[:],
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
