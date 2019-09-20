package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	//监听请求
	listener,err := net.Listen("tcp","0.0.0.0:8000")
	if err != nil{
		fmt.Println("net listen error")
		return
	}
	for {
		conn,err := listener.Accept()
		if err != nil{
			fmt.Println("listen accetp error")
			return
		}
		//handClientRequest(conn)  //这里只能接收一个客户端的请求，如果有别的客户端请求过来，就会等待，因为被这个所阻塞。但是换成下面的方式(go routine)，就不会出现阻塞的问题。
		go handClientRequest(conn)
	}
}

func handClientRequest(conn net.Conn)  {
	defer connectStop(conn)
	for {
		//服务端向客户端发送数据
		_,err := io.WriteString(conn,time.Now().Format("15:04:05\n"))
		if err != nil{
			fmt.Println("server send data error")
			return
		}
		time.Sleep(time.Second*1)
	}
}
func connectStop(conn net.Conn)  {
	fmt.Println("clien exit")
	conn.Close()
}