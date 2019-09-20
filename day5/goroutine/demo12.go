package main

import (
	"fmt"
	"sync"
)

/**
	根据Go语言内存模型规范，对于从无缓冲Channel进行的接收(获取值)，发生在对该Channel进行的发送完成之前(设置值)。
	因此，后台线程<-done接收操作完成之后，main线程的done <- 1发送操作才可能完成（从而退出main、退出程序），
	而此时打印工作已经完成了。
 */
func NoBufferSync()  {
	c1 := make(chan int)

	go func() {
		fmt.Println("hello world")
		<- c1 //在没有收到c1的数据之前，就出于阻塞状态。
	}()
	c1 <- 1
	fmt.Println("main")
}

func HasBufferSync()  {
	done := make(chan int,1)
	go func() {
		fmt.Println("hello world")
		done <- 1
	}()
	<- done
	fmt.Println("main")
}

func MutiSync()  {
	done := make(chan int,10)

	/*for i:=0;i<cap(done) ;i++  {
		go func(i int) {
			fmt.Println("hello world")
			done <- i
		}(i)
	}
	for i := 0;i<cap(done) ;i++  {
		fmt.Println(<-done)
	}*/

	go func() {
		/*for i:=0;i<cap(done) ;i++  {
			go func(i int) {
				fmt.Println("hello world")
				done <- i
			}(i)
		}*/
		/*
		在for循环里面执行的goroutine有可能不会有任何输出，也有可能会出现“panic: send on closed channel”错误。为什么呢?
		情况一：因为创建的多个goroutine是并行，有可能都没有来得及执行，就被主线程执行完毕了，所以就不会有任何输出
		情况二：因为创建的多个goroutine是并行，但是下面的close(done)，都还没有来得及执行完所有的go routine，就执行关闭，就会报错了
		for i:=0;i<10 ;i++  {
			go func(i int) {
				fmt.Println("hello world")
				done <- i
			}(i)
		}
		*/
		//下面的代码就会被执行
		for i:=0;i<10 ;i++  {
			fmt.Println("hello world")
			done <- i
		}
		close(done)
	}()

	for value := range done {
		fmt.Println(value)
	}

	fmt.Println("main")
}

/**
	对于这种要等待N个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用sync.WaitGroup来等待一组事件
 */
func MutilSync2()  {
	var wg sync.WaitGroup

	for i:= 0;i<10 ;i++  {
		wg.Add(1)  //用于增加等待事件的个数，必须确保在后台线程启动之前执行
		go func() {
			fmt.Println("hello world")
			wg.Done()   //表示完成一个事件
		}()
	}
	wg.Wait()  //等待全部的事件完成，在事件完成之前，处于阻塞状态
	fmt.Println("main")
}


func main() {

}
