package znet

import (
	"fmt"
	"net"
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
	// 服务绑定的websocket 端口 (Websocket port the server is bound to)
	WsPort int
	// 服务绑定的websocket 路径 (Websocket path the server is bound to)
	WsPath string
	// 服务绑定的kcp 端口 (kcp port the server is bound to)
	KcpPort int
}

func (s *Server) Start() {
	go startServer(s)
}

func startServer(s *Server) {
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
		go func() {
			buf := make([]byte, 1024)
			for {
				cnt, errRead := conn.Read(buf)
				if errRead != nil {
					fmt.Println("Read err:", errRead)
					continue
				}
				_, errWrite := conn.Write([]byte("echo:" + string(buf[:cnt])))
				if errWrite != nil {
					fmt.Println("Write err:", errWrite)
					continue
				}
			}
		}()
	}
}

// Stop stops the server (停止服务)
func (s *Server) Stop() {
}

// Serve runs the server (运行服务)
func (s *Server) Serve() {
	s.Start()
	select {}
}

func (s *Server) ServerName() string {
	return s.Name
}

func NewServer(name string) *Server {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}

func init() {}
