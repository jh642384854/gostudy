package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
	这个是改良src/dev/day5/goroutine/demo10_sync.Mutex.go的版本
	用互斥锁来保护一个数值型的共享资源，麻烦且效率低下。标准库的sync/atomic包对原子操作提供了丰富的支持。
 */

var sum uint64

func worker2(wg *sync.WaitGroup)  {
	defer wg.Done()
	var i uint64
	for i = 0;i<=100 ;i++  {
		//atomic.AddUint64函数调用保证了total的读取、更新和保存是一个原子操作，因此在多线程中访问也是安全的。
		atomic.AddUint64(&sum,i)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go worker2(&wg)
	go worker2(&wg)
	wg.Wait()

	fmt.Println("sync/atomic:",sum)
}
