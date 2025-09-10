package znet

import (
	"fmt"
	"net"
	"zinx-lib/zconf"
	"zinx-lib/ziface"
	_ "zinx-lib/ziface"
	"zinx-lib/zutils"
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

	msgHandler ziface.IMsgHandle

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
	idWorker, _ := zutils.NewIDWorker(int64(1))
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err:", err)
			continue
		}

		conId, _ := idWorker.NextID()

		connection := NewConnection(conn, uint32(conId), s.msgHandler)
		fmt.Println("New connection:", conId)
		connection.Start()
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

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgID, router)
}

func (s *Server) ServerName() string {
	return s.Name
}

func NewServer(name string) *Server {
	// 初始化消息处理器
	handler := newMsgHandler()

	s := &Server{
		Name:       name,
		IPVersion:  "tcp4",
		IP:         zconf.GlobalObject.Host,
		Port:       zconf.GlobalObject.TCPPort,
		msgHandler: handler,
	}
	return s
}

func init() {}
