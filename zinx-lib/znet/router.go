package znet

import (
	"fmt"
	"io"
	"zinx-lib/ziface"
	"zinx-lib/zpack"
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
	//_, errWrite := conn.Write([]byte("EchoRouter..." + string(data)))

	dataPack := zpack.DataPack{}
	msg, _ := dataPack.Unpack(data)

	data = make([]byte, msg.GetDataLen())

	if _, err := io.ReadFull(conn, data); err != nil {
		fmt.Println("read full data error:", err)
		return
	}

	msg.SetData(data)

	if _, err := conn.Write([]byte("EchoRouter..." + string(data))); err != nil {
		fmt.Println("write data error:", err)
		return
	}
	return
}
