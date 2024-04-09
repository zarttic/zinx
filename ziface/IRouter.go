package ziface

// IRouter 路由抽象接口
type IRouter interface {
	PreHandle(request IRequest)  //在处理conn业务之前的hook方法
	Handle(request IRequest)     //处理conn业务的主方法
	PostHandle(request IRequest) //在处理conn业务之后的hook方法
}
