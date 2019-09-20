package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/utils"
	"zinx/znet"
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
		//3.向服务器发送数据。这里就需要先进行封包操作
		//①、创建一个数据封包、解包的实例对象
		dp := znet.NewDataPack()
		msg := znet.NewMessage(1,[]byte("hello HelloRouter 0.6"))
		//②、进行封包操作，得到一个二进制的数据流
		binaryMsg,err := dp.Pack(msg)
		if err != nil{
			fmt.Println("send message pack error")
			break
		}
		//③、发送数据
		if _,err = conn.Write(binaryMsg); err != nil{
			fmt.Println("客户端发送数据失败")
			break
		}

		//4.接收服务器传递过来的数据。这里就需要解包的操作
		//①、首先解析二进制流的head信息
		headData := make([]byte,dp.GetHeadLen())
		if _,err := io.ReadFull(conn,headData); err != nil{
			fmt.Println("客户端获取服务器发送信息的head错误")
			break
		}
		//②、解包headData
		recvmsg,err := dp.Unpack(headData)
		if err != nil{
			fmt.Println("get message uppack error")
			break
		}
		//③、在从解包的数据中获取实际的消息内容(首先判断返回数据是否有值)
		if recvmsg.GetMessageLen() > 0{
			serverData := make([]byte,recvmsg.GetMessageLen())
			if _,err := io.ReadFull(conn,serverData); err != nil{
				fmt.Println("客户端获取服务器发送信息的body错误")
				break
			}
			//④、正常获取服务端发送的数据，就进行显示处理
			fmt.Printf("接收到服务器发送的消息如下：%v \n",string(serverData))
		}

		time.Sleep(time.Second)
	}

}
