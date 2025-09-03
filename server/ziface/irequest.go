package ziface

type IRequest interface {
	GetConnection() IConnection // Get the connection information of the request(获取请求连接信息)

	GetData() []byte // Get the data of the request message(获取请求消息的数据)
}
