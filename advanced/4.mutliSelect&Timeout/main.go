package main

import (
	"fmt"
	"time"
)

/**
	多路选择器的使用以及超时的应用
 */
func main() {
	//syncRun()
	testAsync()
}

func testAsync() {
	//asyncService()是一个异步的请求，我们可以通过select多路选择器，来进行超时处理
	select {
	case res := <-asyncService():
		fmt.Println(res)
	case <-time.After(time.Millisecond * 100)://如果这里设置的超时时间小于service()里面睡眠的时间，就会执行这个代码，用来控制超时
		fmt.Println("timeout")
	}
	otherTask()
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
