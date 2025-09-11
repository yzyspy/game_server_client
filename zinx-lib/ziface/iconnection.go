package ziface

import "net"

type IConnection interface {
	Start()

	Stop()

	GetTcpConnection() *net.TCPConn

	GetConnId() uint32

	// 这个connection绑定的客户端地址
	RemoteAddr() net.Addr

	//给客户端发送数据
	Send(data []byte) error

	SendMsg(msgID uint32, data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
