package main

import (
	"fmt"
	"time"
)

/**
	CSP VS Actor
	①、和Actor的直接通讯不同，CSP模式则是通过Channel进行通讯的，更松耦合一些。
	②、GO中channel是有容量限制并且独立于处理Groutine，而如Erlang，Actor模式中的mailbox容量是无限的，接收进程也总是被动地处理消息。
 */
func main() {
	//syncRun()
	testAsync()
}

func testAsync() {
	res := asyncService()
	otherTask()
	//从channel中取出数据
	fmt.Println(<-res)
}

/**
	异步服务
 */
func asyncService() chan string {
	//retCh := make(chan string)   //非buffer的channel
	retCh := make(chan string, 1) //带buffer的channel   这里选择的不同，对执行结果也会有一定的影响
	go func() {
		ret := service()
		fmt.Println("return result")
		retCh <- ret
		fmt.Println("service exit")
	}()
	return retCh
}

/**
	下面是同步运行的方式
 */
func syncRun() {
	fmt.Println(service())
	otherTask()
}

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("otherTask working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask is done.")
}
