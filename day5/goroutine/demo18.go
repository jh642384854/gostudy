package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
	下面的代码会展示资源竞争的问题
 */

var (
	//counter 是所有goroutine都要增加其值的变量
	counter2 int
	// wg用来等待程序结束
	wg2 sync.WaitGroup
)

func main() {
	// wg计数加2，表示要等待2个goroutine
	wg2.Add(2)
	//创建两个goroutine
	go incCounter(1)
	go incCounter(2)
	//等待goroutine结束
	wg2.Wait()
	fmt.Println("Final Counter:",counter2)
}

func incCounter(id int)  {
	defer wg2.Done()
	for count := 0; count <2 ;count ++  {
		//捕获counter的值
		value := counter2
		//当前goroutine从线程退出，并放回到队列
		runtime.Gosched()
		//增加本地value的值
		value ++
		//将该值保存会counter
		counter2 = value
	}
}