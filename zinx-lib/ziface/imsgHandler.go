// @Title imsghandler.go
// @Description Provides interfaces for worker startup and handling message business calls
// @Author Aceld - Thu Mar 11 10:32:29 CST 2019
package ziface

// IMsgHandle Abstract layer of message management(消息管理抽象层)
type IMsgHandle interface {
	// Add specific handling logic for messages, msgID supports int and string types
	// (为消息添加具体的处理逻辑, msgID，支持整型，字符串)
	AddRouter(msgID uint32, router IRouter)

	DoMsgHandler(request IRequest)
}
