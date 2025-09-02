package znet

import _ "zinx-server/ziface"

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

}

// Stop stops the server (停止服务)
func (s *Server) Stop() {
}

// Serve runs the server (运行服务)
func (s *Server) Serve() {

}

func (s *Server) ServerName() string {
	return s.Name
}

func init() {}
