package ziface

type IRouter interface {
	PreHandle(request IRequest)  //Hook method before processing conn business(在处理conn业务之前的钩子方法)
	Handle(request IRequest)     //Method for processing conn business(处理conn业务的方法)
	PostHandle(request IRequest) //Hook method after processing conn business(处理conn业务之后的钩子方法)
}
