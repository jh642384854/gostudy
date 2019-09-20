package main

import (
	"fmt"
	"sync"
)

/**
	使用锁机制来实现同步机制
 */

/**
	并发编程的核心概念是同步通信，但是同步的方式却有多种。我们先以大家熟悉的互斥量sync.Mutex来实现同步通信。
	根据文档，我们不能直接对一个未加锁状态的sync.Mutex进行解锁，这会导致运行时异常。下面这种方式并不能保证正常工作：
 */
func lockError()  {
	var lock sync.Mutex

	go func() {
		fmt.Println("hello world")
		lock.Lock()
	}()

	lock.Unlock()
}


/**
	修复的方式是在main函数所在线程中执行两次mu.Lock()，当第二次加锁时会因为锁已经被占用（不是递归锁）而阻塞，
	main函数的阻塞状态驱动后台线程继续向前执行。当后台线程执行到mu.Unlock()时解锁，此时打印工作已经完成了，
	解锁会导致main函数中的第二个mu.Lock()阻塞状态取消，此时后台线程和主线程再没有其它的同步事件参考，
	它们退出的事件将是并发的：在main函数退出导致程序退出时，后台线程可能已经退出了，也可能没有退出。
	虽然无法确定两个线程退出的时间，但是打印工作是可以正确完成的。
 */
func main()  {

	var lock sync.Mutex

	lock.Lock()
	go func() {
		fmt.Println("hello world")
		lock.Unlock()
	}()
	lock.Lock()  //在主线程里面进行两次加锁，第二次就会形成阻塞，然后后台线程就会继续执行。这样goroutine就可以被执行了。

	fmt.Println("main")
}
