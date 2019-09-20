package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//problemCode()

	counter := 0
	var muxt sync.Mutex
	for i := 0;i<5000 ;i++  {
		go func() {
			//在协程内部使用sync.Mutex这个锁机制，解决资源竞争问题。加锁后需要解锁
			defer func() {
				muxt.Unlock()
			}()
			muxt.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("counter:",counter)
	//下面的这方法解决了上面需要通过time.Sleep(time.Second)解决协程执行完毕的问题
	sloveTimeSleep()
}

/**
	下面的这个代码用来解决需要用time.Sleep(time.Second)让协程执行完毕的问题。
	这个需要用到sync.WaitGroup这个内置包来进行处理
 */
func sloveTimeSleep()  {
	counter := 0
	var mutx sync.Mutex
	//声明一个sync.WaitGroup变量
	var wg sync.WaitGroup
	for i := 0;i<5000 ;i++  {
		//向sync.WaitGroup里面添加任务
		wg.Add(1)
		go func() {
			defer func() {
				mutx.Unlock()
			}()
			mutx.Lock()
			counter++
			//表示sync.WaitGroup里面完成了一个任务
			wg.Done()
		}()
	}
	//当sync.WaitGroup处理等待阻塞，直到所有的Add()任务都被Done()才会释放阻塞
	wg.Wait()
	fmt.Println("counter:",counter)
}

/**
	下面的这段代码是有问题的，因为在各个协程里面使用同一个变量，会引发对这个变量的竞争，如果对这个资源没有进行加锁处理的话，最后会导致处理的结果就会出问题。
	所以下面的输出结果并不是每次都是正确的(正确值为5000)，大多数情况下面，结果都是错误的。
 */
func problemCode()  {
	counter := 0
	for i:=0;i<5000 ; i++ {
		go func() {
			counter ++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("counter:",counter)
}