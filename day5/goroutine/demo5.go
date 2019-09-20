package main

import (
	"fmt"
)

/**
	应用实例：
	①、开启一个writeData协程，向管道intChan中写入50个整数
	②、开启一个readData协程，从管道intChan中读取writeData写入的数据
	③、注意：writeData和readData操作的是同一个管道
	④、主线程需要等待writeData和readData协程都完成工作以后才能退出
 */

//写入数据
func WriteData(numChan chan int)  {
	for i:=0;i<50 ;i++  {
		numChan <- i
	}
	//关闭通道
	close(numChan)
}

//读取数据
func ReadData(numChan chan int,exitChan chan bool)  {
	for {
		v,ok := <- numChan
		if !ok{
			break
		}
		fmt.Println("numChan value is :",v)
	}
	//向一个标识完成goroutine的channel中添加数据
	exitChan <- true
	//关闭exitChan
	close(exitChan)
}

func main() {
	numChan := make(chan int,50)
	exitChan := make(chan bool,1)

	go WriteData(numChan)

	go ReadData(numChan,exitChan)

	for {
		_,ok:= <- exitChan
		fmt.Println(ok)
		if ok {
			break
		}
	}

}
