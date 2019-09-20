package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup)  {
	defer wg.Done()  // 操作完成，减少一个计数

	for i := 0;i<=100 ;i++  {
		/**
			在worker的循环中，为了保证total.value += i的原子性，我们通过sync.Mutex加锁和解锁来保证该语句在同一时刻只被一个线程访问。
			对于多线程模型的程序而言，进出临界区前后进行加锁和解锁都是必须的。如果没有锁的保护，total的最终值将由于多线程之间的竞争而可能会不正确。
		 */
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2) // 因为有两个动作，所以增加2个计数
	go worker(&wg)
	go worker(&wg)
	wg.Wait()        // 等待，直到计数为0
	fmt.Println("sync.Mutex:",total.value)
}