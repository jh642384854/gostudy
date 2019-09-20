package main

import (
	"fmt"
	"time"
)

/*
	goroutine的基本示例：

 */

func Test1()  {

	for i := 0;i<=10 ;i++  {
		fmt.Println("Test1() hello world")
		time.Sleep(time.Second)
	}
}

func main() {

	go Test1() //开启一个协程，将会和主线程一起运行，不过，如果这个协程在主线程结束之前还没有运行完毕，也会被终止。

	for i := 0;i<=10 ;i++  {
		fmt.Println("main() hello world")
		time.Sleep(time.Second)
	}
}
