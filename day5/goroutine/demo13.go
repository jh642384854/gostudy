package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/**
	生产者消费者模型
	并发编程中最常见的例子就是生产者消费者模式，该模式主要通过平衡生产线程和消费线程的工作能力来提高程序的整体处理数据的速度。
	简单地说，就是生产者生产一些数据，然后放到成果队列中，同时消费者从成果队列中来取这些数据。这样就让生产消费变成了异步的两个过程。
	当成果队列中没有数据时，消费者就进入饥饿的等待中；而当成果队列中数据已满时，生产者则面临因产品挤压导致CPU被剥夺的下岗问题。
 */

 //定义一个生产者，往channel中写入数据
func Producer(factor int,in chan <- int)  {
	for i:=0; ;i++  {
		in <- i*factor
	}
}
//定义一个消费者，从channel中读取数据
func Consumer(out <- chan int)  {
	for value := range out {
		 fmt.Println(value)
	}
}

func main() {
	nums := make(chan int,64)

	go Producer(3,nums)
	go Producer(5,nums)

	go Consumer(nums)

	//设定5秒钟以后退出主程序
	//time.Sleep(5*time.Second)
	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}