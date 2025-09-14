package znet

import (
	"fmt"
	"zinx-lib/ziface"
)

type MsgHandler struct {
	//key: msgId  value: IRouter 一个msgId一个处理逻辑
	Apis map[uint32]ziface.IRouter

	// A message queue for workers to take tasks
	// (Worker负责取任务的消息队列)
	TaskQueue []chan ziface.IRequest

	WorkerPoolSize uint32
}

func newMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: 8,
		TaskQueue:      make([]chan ziface.IRequest, 8),
	}
}

func (m *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	m.Apis[msgID] = router
}

func (m *MsgHandler) DoMsgHandler(request ziface.IRequest, workerID int) {
	fmt.Println("workerID:", workerID, "msgId:", request.GetMsgId())
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

func (m *MsgHandler) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	// Continuously wait for messages in the queue
	// (不断地等待队列中的消息)
	for {
		select {
		// If there is a message, take out the Request from the queue and execute the bound business method
		// (有消息则取出队列的Request，并执行绑定的业务方法)
		case request, ok := <-taskQueue:
			if !ok {
				return
			}
			switch req := request.(type) {
			case ziface.IRequest: // Client message request
				m.DoMsgHandler(req, workerID)
			}
		}
	}
}

// StartWorkerPool starts the worker pool
func (mh *MsgHandler) StartWorkerPool() {
	// Iterate through the required number of workers and start them one by one
	// (遍历需要启动worker的数量，依此启动)
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// A worker is started
		// Allocate space for the corresponding task queue for the current worker
		// (给当前worker对应的任务队列开辟空间)
		mh.TaskQueue[i] = make(chan ziface.IRequest, 1000)

		// Start the current worker, blocking and waiting for messages to be passed in the corresponding task queue
		// (启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来)
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}
