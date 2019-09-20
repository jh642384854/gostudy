package main

import (
	"fmt"
	"sync/atomic"
	"time"
	"sync"
)

var (
	shutdown int64
	wg4 sync.WaitGroup
)
/**
	atomic.StoreInt64()和atomic.LoadInt64()的使用
 */

func main() {

	wg4.Add(2)
	go doWork("A")
	go doWork("B")
	//给定goroutine执行的时间
	time.Sleep(1*time.Second)
	fmt.Println("shutdown now")
	atomic.StoreInt64(&shutdown,1)
	//等待goroutine结束
	wg4.Wait()
}

func doWork(str string)  {
	defer wg4.Done()
	for  {
		fmt.Printf("Doing %s Work \n",str)
		time.Sleep(205*time.Millisecond)

		//要停止工作了吗？
		if atomic.LoadInt64(&shutdown) == 1{
			fmt.Printf("Doing %s Shutdown \n",str)
			break
		}
	}
}