package ziface

// Defines the server interface
type IServer interface {
	Start() // Start the server method(启动服务器方法)
	Stop()  // Stop the server method (停止服务器方法)
	Serve() // Start the business service method(开启业务服务方法)
	// Get the server name (获取服务器名称)
	ServerName() string
}
