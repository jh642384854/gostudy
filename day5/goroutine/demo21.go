package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
	互斥锁解决资源竞争问题
 */

var (
	//counter 是所有goroutine都要增加其值的变量
	counter5 int64
	// wg用来等待程序结束
	wg5 sync.WaitGroup
	// mutex用来定义一段代码临界区
	mutex sync.Mutex
)

func main() {
	// wg计数加2，表示要等待2个goroutine
	wg5.Add(2)
	//创建两个goroutine
	go incCounter3(1)
	go incCounter3(2)
	//等待goroutine结束
	wg5.Wait()
	fmt.Println("Final Counter:",counter5)
}

func incCounter3(id int)  {
	defer wg5.Done()
	for count := 0; count <2 ;count ++  {
		//同一时刻只允许一个 goroutine 进入  这是一个临界区
		mutex.Lock()
		{
			//捕获counter的值
			value := counter5
			//当前goroutine从线程退出，并放回到队列
			runtime.Gosched()
			//增加本地value的值
			value ++
			//将该值保存会counter
			counter5 = value
		}
		//释放锁，允许其他正在等待的 goroutine
		mutex.Unlock()
	}
}