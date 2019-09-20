package znet

import (
	"fmt"
	"zinx/ziface"
)

type PingRouter struct {
	BaseRouter
}

//处理实际业务之前的方法
/*func (this *PingRouter) PrevHandler(request ziface.IRequest) {
	fmt.Println("call prev PrevHandler")
	_, err := request.GetConnect().GetTCPConnection().Write([]byte("before ping"))
	if err != nil {
		fmt.Println("call prev handler error")
	}

}*/

//处理实际业务的主方法
func (this *PingRouter) Handler(request ziface.IRequest) {
	/*
	fmt.Println("call handler")
	_, err := request.GetConnect().GetTCPConnection().Write([]byte("ping... ping... ping..."))
	if err != nil {
		fmt.Println("call  handler error")
	}
	*/
	fmt.Println("call handler")
	//①、获取客户端发送给服务端的数据
	fmt.Printf("Message id :%d,Message Content:%v \n",request.GetDataID(),string(request.GetData()))
	//②、服务端向客户端发生数据
	err := request.GetConnect().SendData(0,[]byte("ping ... ping ..."))
	if err != nil{
		fmt.Println("server send data errror")
	}
}

//处理实际业务之后的方法
/*func (this *PingRouter) PostHandler(request ziface.IRequest) {
	fmt.Println("call afer PostHandler")
	_, err := request.GetConnect().GetTCPConnection().Write([]byte("afer ping"))
	if err != nil {
		fmt.Println("call afer handler error")
	}
}*/
