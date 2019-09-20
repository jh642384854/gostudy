package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/**
	下面的代码会通过原子操作来解决资源竞争的问题
 */

var (
	//counter 是所有goroutine都要增加其值的变量
	counter3 int64
	// wg用来等待程序结束
	wg3 sync.WaitGroup
)

func main() {
	// wg计数加2，表示要等待2个goroutine
	wg3.Add(2)
	//创建两个goroutine
	go incCounter2(1)
	go incCounter2(2)
	//等待goroutine结束
	wg3.Wait()
	fmt.Println("Final Counter:",counter3)
}

func incCounter2(id int)  {
	defer wg3.Done()
	for count := 0; count <2 ;count ++  {
		//安全地对counter3加1
		atomic.AddInt64(&counter3,1)
		runtime.Gosched()
	}
}