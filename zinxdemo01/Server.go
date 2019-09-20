package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

/**
	基于zinx框架来开发   这是服务端应用
	可借鉴的框架代码：
	https://goframe.org/container/gset/index
 */

func EventConnectionBegin(connection ziface.IConnection)  {
	fmt.Println("EventConnectionBegin：在客户端建立链接后立即被执行的方法")
	connection.SendData(202,[]byte("广播消息"))
	connection.SetProperty("Name","zinx study is very good")
	connection.SetProperty("Author","aceld")
}

func EventConnectionStop(connection ziface.IConnection)  {
	fmt.Println("EventConnectionStop：在客户端断开链接前立即被执行的方法,当前链接ID：",connection.GetConnectionID())
	if val,ok := connection.GetProperty("Name");ok == nil{
		fmt.Println("Name:",val)
	}
	if val,ok := connection.GetProperty("Author");ok == nil{
		fmt.Println("Author:",val)
	}
}



func main() {
	//1.创建一个Server句柄，使用zinx的api
	s := znet.NewServer("[ zinx v0.9 ]")

	//为当前链接链接绑定前置和后置处理方法(就是当前链接建立后，和当前链接被断开前马上执行的方法或事件)
	s.SetOnConnectStart(EventConnectionBegin)
	s.SetOnConnectStop(EventConnectionStop)

	s.AddRouter(0,&znet.PingRouter{})
	s.AddRouter(1,&znet.HelloRouter{})
	//2.启动Server
	s.Run()
}
