package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Test1()  {
	conn,err := net.Dial("tcp","0.0.0.0:8000")
	if err != nil{
		fmt.Println("client connect error")
		return
	}
	//接收服务端的数据
	defer conn.Close()
	mustCopy(os.Stdout,conn)
}

func main() {
	conn,err := net.Dial("tcp","0.0.0.0:8000")
	if err != nil{
		log.Fatal(err)
		return
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout,conn)
		log.Println("donw")
		done <- struct{}{}
	}()
	mustCopy(conn,os.Stdin)
	conn.Close()
	<- done  //这里会阻塞等待相应的goroutine处理完后，在执行
}
func mustCopy(dst io.Writer,src io.Reader)  {
	if _,err := io.Copy(dst,src); err != nil{
		log.Fatal(err)
	}
}