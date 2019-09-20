package ziface

/**
	基础路由的定义
	路由里的数据的，都是基于IRouter来处理的。因为IRouter里面包含了客户端传递的数据了
	需要包含处理业务的方法：可以包含处理实际业务之前、包含处理实际业务之后，以及真正实现业务的方法
 */

type IRouter interface {
	//处理实际业务之前的方法
	PrevHandler(request IRequest)
	//处理实际业务的主方法
	Handler(request IRequest)
	//处理实际业务之后的方法
	PostHandler(request IRequest)
}