package ziface

/**
	请求处理接口
	将客户端请求的TCP链接和数据绑定在一起，就定义在这包里面。
 */

 type IRequest interface {
 	//得到当前的链接，这里的链接，就用我们创建好的IConnection模块
 	GetConnect() IConnection

 	//得到当前的消息数据
	GetData() []byte

 	//得到当前的消息的ID
	GetDataID() uint32
 }
