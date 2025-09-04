package znet

import (
	"fmt"
	"zinx-server/ziface"
)

type BaseRouter struct {
}

func (b *BaseRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("BaseRouter PreHandle")
}

func (b *BaseRouter) Handle(request ziface.IRequest) {
	fmt.Println("BaseRouter Handle")
}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("BaseRouter PostHandle")
}

type EchoRouter struct {
	BaseRouter
}

func (e *EchoRouter) Handle(request ziface.IRequest) {
	fmt.Println("EchoRouter Handle")
	conn := request.GetConnection().GetTcpConnection()
	data := request.GetData()
	// 向客户端发送数据
	_, errWrite := conn.Write([]byte("EchoRouter to client:" + string(data)))
	if errWrite != nil {
		fmt.Println("Write err:", errWrite)
		return
	}
}
