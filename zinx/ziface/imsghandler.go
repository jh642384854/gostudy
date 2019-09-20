package ziface

/**
	多路由实现模块：用来管理多个消息处理。
	根据客户端发送的消息不一样，使用不同的路由来处理
 */

type IMsghandler interface {
	//执行指定消息的路由信息
	DoMsgHandler(request IRequest) error
	//添加路由到消息管理器中
	AddMsgHandler(msgid uint32, router IRouter) error
	//启动工作池
	StartWorkPool()
	//将请求添加到任务队列中
	AddRequestToTaskQueue(workRequest IRequest)
}
