package znet

import (
	"fmt"
	"zinx-lib/ziface"
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
	conn := request.GetConnection()
	data := request.GetData()

	conn.SendMsg(101, data)
}
