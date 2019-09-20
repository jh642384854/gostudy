package main

import (
	"fmt"
	"net"
	"time"
	"zinx/utils"
)

/**
	模拟客户端
 */
func main() {
	fmt.Println("client start ......")
	//1.直接连接远程服务器，得到一个conn的连接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", utils.GloabConfigObj.Host, utils.GloabConfigObj.Port))
	if err != nil {
		fmt.Println("客户端连接建立失败")
		return
	}
	//2.通过建立的连接来调用Write()方法写数据
	for {
		//3.向服务器发送数据
		_, err := conn.Write([]byte("hello zinx 0.2"))
		if err != nil {
			fmt.Println("客户端发送数据失败")
			return
		}
		//4.接收服务器传递过来的数据
		buf := make([]byte, 512)
		contentLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端接收数据失败")
			return
		}
		fmt.Printf("服务端数据：%s,长度：%d \n", buf, contentLen)

		time.Sleep(time.Second)
	}

}
