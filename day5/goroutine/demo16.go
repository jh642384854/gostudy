package main

import (
	"fmt"
	"runtime"
	"sync"
)
func main() {

	// 分配一个逻辑处理器给调度器使用
	//runtime.GOMAXPROCS(1)
	// 分配2个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)
	//runtime.GOMAXPROCS(runtime.NumCPU())
	// wg用来等待程序完成。计数器加2，标识要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")
	//声明一个匿名函数，并创建一个goroutine
	go func() {
		//在函数退出时调用Done来通知main函数工程已经完成
		defer wg.Done()
		for count := 0; count < 3 ; count++  {
			for char := 'a'; char < 'a'+26 ; char++  {
				fmt.Printf("%c",char)
			}
		}
	}()

	go func() {
		//在函数退出时调用Done来通知main函数工程已经完成
		defer wg.Done()
		for count := 0; count < 3 ; count++  {
			for char := 'A'; char < 'A'+26 ; char++  {
				fmt.Printf("%c",char)
			}
		}
	}()

	fmt.Println("\n Waiting To Finish")
	//等待上面的gorutine执行完毕
	wg.Wait()
	fmt.Println("\n Terminating Program")
}