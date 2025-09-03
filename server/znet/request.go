package znet

import (
	"zinx-server/ziface"
)

type Request struct {
	conn ziface.IConnection // the connection which has been established with the client(已经和客户端建立好的连接)
	msg  ziface.IMessage    // the request data sent by the client(客户端请求的数据)

	data []byte // the response data to the client(响应给客户端的数据)
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
