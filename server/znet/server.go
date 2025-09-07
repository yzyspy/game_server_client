package znet

import (
	"fmt"
	"math/rand"
	"net"
	"time"
	"zinx-server/zconf"
	"zinx-server/ziface"
	_ "zinx-server/ziface"
)

// Server interface implementation, defines a Server service class
// (接口实现，定义一个Server服务类)
type Server struct {
	// Name of the server (服务器的名称)
	Name string
	//tcp4 or other
	IPVersion string
	// IP version (e.g. "tcp4") - 服务绑定的IP地址
	IP string
	// IP address the server is bound to (服务绑定的端口)
	Port int

	Router ziface.IRouter

	// 服务绑定的websocket 端口 (Websocket port the server is bound to)
	WsPort int
	// 服务绑定的websocket 路径 (Websocket path the server is bound to)
	WsPath string
	// 服务绑定的kcp 端口 (kcp port the server is bound to)
	KcpPort int
}

func (s *Server) Start() {
	go s.startServer()
}

func (s *Server) startServer() {
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("ResolveTCPAddr err:", err)
		return
	}

	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("ListenTCP err:", err)
		return
	}
	fmt.Println("Server is running on", listener.Addr().String())
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err:", err)
			continue
		}
		rand.Seed(time.Now().UnixNano())
		conId := rand.Uint32()
		connection := NewConnection(conn, conId, s.Router)
		fmt.Println("New connection:", conId)
		connection.Start()
	}
}

func echoFunc(conn *net.TCPConn, data []byte, cnt int) error {
	_, errWrite := conn.Write([]byte("echo to client:" + string(data)))
	if errWrite != nil {
		fmt.Println("Write err:", errWrite)
		return errWrite
	}
	return nil
}

// Stop stops the server (停止服务)
func (s *Server) Stop() {
}

// Serve runs the server (运行服务)
func (s *Server) Serve() {
	s.Start()
	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}

func (s *Server) ServerName() string {
	return s.Name
}

func NewServer(name string) *Server {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        zconf.GlobalObject.Host,
		Port:      zconf.GlobalObject.TCPPort,
		Router:    nil,
	}
	return s
}

func init() {}
