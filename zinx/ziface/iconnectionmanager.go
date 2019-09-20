package ziface

/**
	客户端链接管理模块
	用于增、删、查用户与服务端建立的链接
 */
type IConnectionManager interface {
	//添加链接
	Add(conn IConnection)
	//根据链接ID来删除链接
	Remove(connid uint32)
	//根据链接ID来获取链接
	Get(connid uint32) (IConnection,error)
	//得到当前链接管理器的所有链接总数。这个用来控制流量，设置当前允许最大链接数
	Len() int
	//清除所有链接
	ClearAll()
}