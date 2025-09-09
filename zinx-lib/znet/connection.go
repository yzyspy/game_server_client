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
