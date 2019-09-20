package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

//iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	//定义服务的一些基本属性
	Name      string //服务器的名称
	IPVersion string //服务器绑定的IP版本+
	IPAddress string //服务器绑定的IP地址
	Port      int    //服务器监听的端口
	//Router ziface.IRouter //添加router成员
	MsgHandler ziface.IMsghandler
	//客户端链接管理模块
	ConnMgr    ziface.IConnectionManager
	//客户端建立链接后首先触发的函数
	OnConnectStart func(connection ziface.IConnection)
	//客户端断开链接前触发的函数
	OnConnectStop func(connection ziface.IConnection)
}

//定义一个实现src/zinx/ziface/iconnection.go中的HandleFunc方法
func CallBackToClient(conn *net.TCPConn, buf []byte, contlen int) error {
	if _, err := conn.Write(buf[:contlen]); err != nil {
		return errors.New("读取客户端传递的数据出错")
	}
	return nil
}

//启动服务器
func (this *Server) Start() {
	fmt.Printf("zinx server start success,severname :%s,port:%d,ip:%s", utils.GloabConfigObj.ServerName, utils.GloabConfigObj.Port,utils.GloabConfigObj.Host)
	go func() {
		//开启任务任务工作池
		this.MsgHandler.StartWorkPool()

		//1.获取一个TCP的addr
		tcpAddres, err := net.ResolveTCPAddr(this.IPVersion, fmt.Sprintf("%s:%d", this.IPAddress, this.Port))
		if err != nil {
			fmt.Println("服务器连接失败")
			return
		}
		//2.监听服务器的地址
		tcpListen, err := net.ListenTCP(this.IPVersion, tcpAddres)
		if err != nil {
			fmt.Println("服务器监听失败")
			return
		}
		fmt.Println("start zinx server succ")
		//3.阻塞的等待客户端连接，处理客户端链接业务(读写)
		var connid uint32 = 0
		for {
			tcpConn, err := tcpListen.AcceptTCP()
			if err != nil {
				fmt.Println("接收客户端请求失败")
				continue
			}

			//判断当前链接是否已经大于系统设置的最大链接数
			if this.ConnMgr.Len() > utils.GloabConfigObj.MaxConnectionSize{
				fmt.Printf("当前请求过多，请您稍后重试,当前允许最大链接数是：%d，而当前有 %d 链接数 \n",utils.GloabConfigObj.MaxConnectionSize,this.ConnMgr.Len())
				tcpConn.Close()
				continue
			}

			//将处理新连接的业务方法和tcpConn进行绑定，得到我们的链接模块
			dealConn := NewConnection(this,tcpConn, connid, this.MsgHandler)
			connid ++
			//启动当前的链接业务
			go dealConn.Start()
		}
	}()
}

//停止服务器
func (this *Server) Stop() {

}

//运行服务器
func (this *Server) Run() {
	//启动服务器功能
	this.Start()

	//阻塞状态
	select {}
}

func (this *Server) AddRouter(msgid uint32, router ziface.IRouter) {
	//this.Router = router
	this.MsgHandler.AddMsgHandler(msgid, router)
}

//初始化Server模块的方法
/**
	注意这里的返回值是ziface.IServer这个接口对象。为什么是这个接口对象呢？
	因为我们还可以依据这个接口来实现很多别的服务，比如这里我们实现的是基于TCP协议的。
	我们当然还可以基于ziface.IServer这个接口对象来实现一个UDP协议的服务。
 */
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GloabConfigObj.ServerName,
		IPVersion:  "tcp4",
		IPAddress:  utils.GloabConfigObj.Host,
		Port:       utils.GloabConfigObj.Port,
		MsgHandler: NewMsgHandler(),
		ConnMgr:    NewConnectionManager(),
	}
	return s
}

//得到ConnectionManager对象
func (this *Server) GetConnManager() ziface.IConnectionManager {
	return this.ConnMgr
}

//注册客户端建立链接后马上调用的事件
func (this *Server) SetOnConnectStart(hookFun func(connection ziface.IConnection))  {
	this.OnConnectStart = hookFun
}
//注册客户端断开链接后马上调用的事件
func (this *Server) SetOnConnectStop(hookFun func(connection ziface.IConnection))  {
	this.OnConnectStop = hookFun
}
//以上是绑定客户端建立和释放链接前置和后置方法，下面就要定义如何调用这两个方法
func (this *Server) CallOnConnectStart(connection ziface.IConnection)  {
	if this.OnConnectStart != nil{
		this.OnConnectStart(connection)
	}
}
func (this *Server) CallOnConnectStop(connection ziface.IConnection)  {
	if this.OnConnectStop != nil{
		this.OnConnectStop(connection)
	}
}
