package ziface

import "net"

/**
	链接封装
 */

type IConnection interface {
	//1.启动链接，让当前链接准备开始工作
	Start()

	//2.停止链接，结束当前的链接工作
	Stop()

	//3.获取当前链接绑定的socket conn
	GetTCPConnection() *net.TCPConn

	//4.获取当前链接模块的链接ID
	GetConnectionID() uint32

	//5.获取远程客户端的TCP状态信息，比如IP和端口等等
	GetRemoteAddr() net.Addr

	//6.发送数据，将数据发送给远程的客户端
	//Send(data []byte) error
	SendData(messageid uint32,data []byte) error

	//为当前链接操作属性定义的一些方法
	//①、设置属性
	SetProperty(key string,val interface{})
	//②、获取属性
	GetProperty(key string)(interface{},error)
	//③、删除属性
	RemoveProperty(key string)
}
//定义一个处理链接业务的方法。这里定义的类型是一个func,这里的方法里面只有参数类型，并不需要形参
/**
	参数1：是TCPConn对象
	参数2：是客户端发送给服务端的数据
	参数3：发送数据的长度
 */
type HandleFunc func(*net.TCPConn,[]byte,int) error