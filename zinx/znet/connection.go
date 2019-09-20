package znet

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"net"
	"sync"
	"zinx/utils"
	"zinx/ziface"
)

type Connection struct {
	//定义当前链接所绑定的服务器
	Server ziface.IServer

	//当前链接的socket TCP套接字
	Conn *net.TCPConn

	//链接ID
	ConnID uint32

	//当前链接的状态
	IsClosed bool

	//当前链接所绑定的处理业务方法API
	//handAPI ziface.HandleFunc  //在v0.3里面废除，使用router来定义业务处理

	//告知当前链接已经退出或停止  channel （由Reader告知Writer退出）
	ExitChan chan bool

	//无缓冲的管道，用于读、写go routine之间的消息通信
	MsgChan chan []byte

	//该链接处理的router
	//Router ziface.IRouter
	//当前链接的实际处理handler
	MsgHandler ziface.IMsghandler

	//为当前链接设置的一些自定义属性
	propertyMap map[string]interface{}
	//操作当前链接属性时候添加锁机制
	propertyLock sync.RWMutex
}

//初始化Connection方法
/*
V0.2版本里面使用的是HandleFunc回调的方式来处理响应，v0.3里面使用的router方式俩处理
func NewConnection(conn *net.TCPConn,connID uint32,callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:conn,
		ConnID:connID,
		handAPI:callback_api,
		IsClosed:false,
		ExitChan:make(chan bool, 1),
	}
	return  c
}
*/
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsghandler) *Connection {
	c := &Connection{
		Server:      server,
		Conn:        conn,
		ConnID:      connID,
		MsgHandler:  msgHandler,
		IsClosed:    false,
		ExitChan:    make(chan bool, 1),
		MsgChan:     make(chan []byte),
		propertyMap: make(map[string]interface{}),
	}
	//将当前链接添加到链接管理器中
	server.GetConnManager().Add(c)
	return c
}

//从服务端读取数据的方法
func (this *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is Running......]")

	defer this.Stop()
	defer fmt.Println("Conn：", this.ConnID, " Reader is exit,remote add is ", this.GetRemoteAddr().String())

	//循环读取(阻塞)
	for {
		/*
				版本1 获取客户端传递的数据后，该如何处理呢？这里就要调用我们定义的自定义handapi来处理了
				使用了路由后，就不需要在使用单独定义的handlerApi的方式了
				if err := this.handAPI(this.Conn,buf,conlength); err != nil{
					fmt.Println("ConnID:",this.ConnID," handler is error")
					break //如果调用用户自定义的handle出错，就退出请求
				}
		*/
		/*
				版本2 没有使用消息封装类进行消息处理的实现过程
				buf := make([]byte,512)
				_,err := this.Conn.Read(buf)
				if err != nil{
					fmt.Println("recv data err",err)
					continue  //如果本次读取数据出错，则跳过，等待下一次客户端传递数据继续读取
				}
		*/
		//版本3 使用实现的消息封装处理类来对消息进行封包和解包处理
		//①、首先创建一个数据封包、解包的类
		dp := NewDataPack()
		//②、从二进制流中读取head信息
		headData := make([]byte, dp.GetHeadLen())
		if _, err := this.Conn.Read(headData); err != nil {
			fmt.Println("get message head error")
			break
		}
		//③、将得到的head进行拆包，得到MsgID和datalen
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack data error")
			break
		}
		//④、依据上面读取的datalen，再次从conn里面读取数据，得到的就是实际消息内容，并将消息内容放入到Message对象中
		if msg.GetMessageLen() > 0 {
			body := make([]byte, msg.GetMessageLen())
			if _, err := this.Conn.Read(body); err != nil {
				fmt.Println("get message body error")
				break
			}
			msg.SetMessageData(body)
		}
		request := &Request{
			Conn:    this,
			msgData: msg,
		}
		//判断是否开启了工作池处理方式来处理请求
		if utils.GloabConfigObj.MaxWorkPoolSize > 0 {
			this.MsgHandler.AddRequestToTaskQueue(request)
		} else {
			go this.MsgHandler.DoMsgHandler(request)
		}

	}
}

/**
	主要用来实现写消息。专门发送消息给客户端
 */
func (this *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is Running......]")
	defer fmt.Println("Conn：", this.ConnID, " Writer is exit,remote add is ", this.GetRemoteAddr().String())

	//循环从MsgChan读取数据，然后发给客户端
	for {
		select {
		//如果有消息在MsgChan这个管道里面，则就读取消息并发送
		case data := <-this.MsgChan:
			if _, err := this.Conn.Write(data); err != nil {
				fmt.Println("send data error", err)
				return
			}
		//如果接收到了链接退出的信号，则也表示当前给客户端返回数据也退出
		case <-this.ExitChan:
			return
		}
	}

}

//实现zinx/ziface的IConnection这个方法
//1.启动链接，让当前链接准备开始工作
func (this *Connection) Start() {
	fmt.Println("Conn Start ... ConnID = ", this.ConnID)
	go this.StartReader()
	go this.StartWriter()
	//在建立链接之后调用链接的钩子函数
	this.Server.CallOnConnectStart(this)
}

//2.停止链接，结束当前的链接工作
func (this *Connection) Stop() {
	fmt.Println("conn stop .. ConnID:", this.ConnID)

	//判断是否已经关闭
	if this.IsClosed {
		return
	}
	this.IsClosed = true
	//告知Writer退出
	this.ExitChan <- true
	//在当前链接断开之前调用链接断开的钩子函数
	this.Server.CallOnConnectStop(this)
	//关闭socket链接
	this.Conn.Close()
	//关闭链接
	this.Server.GetConnManager().Remove(this.ConnID)

	//关闭管道，进行资源回收
	close(this.ExitChan)
	close(this.MsgChan)
}

//3.获取当前链接绑定的socket conn
func (this *Connection) GetTCPConnection() *net.TCPConn {
	return this.Conn
}

//4.获取当前链接模块的链接ID
func (this *Connection) GetConnectionID() uint32 {
	return this.ConnID
}

//5.获取远程客户端的TCP状态信息，比如IP和端口等等
func (this *Connection) GetRemoteAddr() net.Addr {
	return this.Conn.RemoteAddr()
}

//6.发送数据，将数据发送给远程的客户端，我们需要将发生的数据进行封包处理，然后在发送
func (this *Connection) SendData(mid uint32, data []byte) error {
	//①、如果当前连接已经被关闭，返回一个异常信息
	if this.IsClosed == true {
		return errors.New("Connection close when send message")
	}
	//②、要对数据进行封包处理，就需要新创建一个封包对象
	dp := NewDataPack()
	var sendData []byte
	switch mid {
	case 0:
		sendData = []byte("hello PingRouter")
	case 1:
		sendData = []byte("hello HelloRouter")
	default:
		sendData = []byte("hello zinx study")
	}
	msg := NewMessage(0, sendData)
	//通过封包操作，就会得到二进制的消息数据
	binaryMsg, err := dp.Pack(msg);
	if err != nil {
		return errors.New("data pack error")
	}
	//③、发生数据给客户端
	/*if _, err := this.Conn.Write(binaryMsg); err != nil {
		return errors.New("server send data error")
	}*/
	//④、使用channel后，就会把消息发送到这个channel中
	//服务端写的数据先放在一个管道中，比如这里的MsgChan，好处就在于这里做了读和写的分离，我们可以对这个数据进行二次处理，比如数据转换、数据合法性校验等等，这样就会很灵活
	this.MsgChan <- binaryMsg
	return nil
}

//7.为当前链接操作属性定义的一些方法
//①、设置属性
func (this *Connection) SetProperty(key string,val interface{}){
	//开启锁保护机制
	this.propertyLock.Lock()
	defer this.propertyLock.Unlock()
	this.propertyMap[key] = val
}
//②、获取属性
func (this *Connection) GetProperty(key string)(interface{},error){
	//开启锁保护机制
	this.propertyLock.RLock()
	defer this.propertyLock.RUnlock()
	if val,ok := this.propertyMap[key];ok{
		return val,nil
	}else{
		return nil,errors.New(fmt.Sprintf("没有找到key为%v的属性值",key))
	}
}
//③、删除属性
func (this *Connection) RemoveProperty(key string){
	//开启锁保护机制
	this.propertyLock.Lock()
	defer this.propertyLock.Unlock()
	delete(this.propertyMap,key)
}