package main

import (
	"bufio"
	"fmt"
	"net"
)
/**
	使用不同的语言编写的socket程序是可以互通的。
	即比如用go写的socket服务端，使用php写的客户端来连接，这样并没有什么问题。
	反过来，用php写的服务端，用go写的客户端连连接，一样可以正常
 */
func main() {
	listen,err := net.Listen("tcp",":20000")
	if err != nil{
		fmt.Println("listen failed ,err:",err)
		return
	}
	for  {
		//建立链接
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println("acceptd failed,err:",err)
			return
		}
		//启动一个协程去处理连接
		go process(conn)
	}
}

func process(conn net.Conn)  {
	for  {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read(buf[:])
		if err != nil{
			fmt.Println("read from client faild,err:",err)
			break
		}
		receiverStr := string(buf[:])
		fmt.Println("receive client send data:",receiverStr,",length :",n)
		conn.Write([]byte(receiverStr))
	}
	defer conn.Close()
}