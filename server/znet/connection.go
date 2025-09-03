package znet

import (
	"fmt"
	"net"
	"zinx-server/ziface"
)

type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	FuncApi ziface.HandleFunc

	ExitChann chan bool
}

func (c *Connection) Start() {
	go c.StartRead()
}

func (c *Connection) StartRead() {
	buf := make([]byte, 1024)
	for {
		cnt, errRead := c.Conn.Read(buf)
		if errRead != nil {
			fmt.Println("Read err:", errRead)
			continue
		}
		c.FuncApi(c.Conn, buf[:cnt], cnt)
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

func NewConnection(conn *net.TCPConn, connID uint32, funcApi ziface.HandleFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		FuncApi:   funcApi,
		isClosed:  false,
		ExitChann: make(chan bool, 1),
	}
}
