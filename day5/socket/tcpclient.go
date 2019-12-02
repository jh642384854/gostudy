package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn,err := net.Dial("tcp",":30000")
	if err != nil{
		fmt.Println("client connect server err:",err)
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for  {
		//读取用户输入
		input,_ := inputReader.ReadString('\n')
		inputStr := strings.Trim(input,"\r\n")
		//如果输入Q就退出
		if strings.ToUpper(inputStr) == "Q"{
			return
		}
		//向服务端发送数据
		_,err := conn.Write([]byte(inputStr))
		if err != nil{
			fmt.Println("client write content faild",err)
			return
		}
		//读取服务端发送过来的数据
		buf := [512]byte{}
		_,err = conn.Read(buf[:])
		if err != nil{
			fmt.Println("received server content faild,err:",err)
			return
		}
		fmt.Println(string(buf[:]))
	}
}