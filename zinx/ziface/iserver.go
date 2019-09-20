package ziface

type IServer interface {
	//启动服务
	Start()
	//停止服务
	Stop()
	//运行服务
	Run()
	//给当前的服务注册一个路由方法，供客户端的链接使用
	//AddRouter(router IRouter)
	//多路由实现
	AddRouter(msgid uint32,router IRouter)
	//得到当前链接管理器
	GetConnManager() IConnectionManager
	SetOnConnectStart(hookFun func(connection IConnection))
	SetOnConnectStop(hookFun func(connection IConnection))
	CallOnConnectStart(connection IConnection)
	CallOnConnectStop(connection IConnection)
}