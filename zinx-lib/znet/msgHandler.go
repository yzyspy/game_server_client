package znet

import (
	"fmt"
	"zinx-lib/ziface"
)

type MsgHandler struct {
	//key: msgId  value: IRouter 一个msgId一个处理逻辑
	Apis map[uint32]ziface.IRouter
}

func newMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

func (m *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	m.Apis[msgID] = router
}

func (m *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	//根据msgId找到对应的处理逻辑
	router := m.Apis[request.GetMsgId()]
	if router == nil {
		fmt.Println("router is nil.................")
		return
	}
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}
