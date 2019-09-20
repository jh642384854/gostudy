package main

import (
	"fmt"
	"math/rand"
	"time"
)
//生成数据写入到channel中
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for  {
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
			//time.Sleep(1*time.Millisecond)
			out <-i
			i++
		}
	}()
	return out
}

func doworker(id int,c chan int)  {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d \n",id,n)
	}
}

func createWorker(id int) chan <- int {
	c := make(chan int)
	go doworker(id,c)
	return c
}

/**
	select case结构适用于多任务并行操作：
	会根据case条件的成立来执行具体那个任务内容。

 */
func main() {
	var c1,c2 = generator(),generator()
	var worker = createWorker(0)
	var values []int //将生成的数据保存到这个数组中，以免因为程序处理过快，漏掉部分数据

	tm := time.After(10*time.Second)  //定义结束程序的定时器，这里设置的是10s后就结束程序运行
	tick := time.Tick(time.Second)    //定义获取values长度的定时器，这里设置的是每隔1s就获取一下values长度
	for {
		var activeWorker chan <- int
		var activeValue int //当前要被处理的数据
		//当接收到了数据后，开始执行任务
		if len(values) > 0{
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <- c1:
			values = append(values,n)
		case n := <- c2:
			values = append(values,n)
		case activeWorker <- activeValue:  //当要被处理的数据有值的时候，将里面的数据提出一个
			values = values[1:]
		case <- time.Tick(800*time.Millisecond):  //这里设置的是每隔800毫秒来运行一次，只要其他操作超过800毫秒，就执行这里的代码
			fmt.Println("time out")
		case <- tick:
			fmt.Println("queue len :",len(values))
		case <- tm:
			fmt.Println("bye")
			return
		}
	}
}


/**
	上面的generator()函数在生成数据的时候，会睡眠一会，但是我们执行下面的函数的时候，可能只会看到nothing，其实并不是的，
	只不过是程序执行的太快，看不到其他数字的生成
 */
func selectCaseV1()  {
	var c1,c2 = generator(),generator()

	for {
		select {
		case n := <- c1:
			fmt.Println(n)
		case n := <- c2:
			fmt.Println(n)
		default:
			fmt.Println("nothing")
		}
	}
}
/**
	使用独立的worker来处理
 */
func selectCaseV2()  {
	var c1,c2 = generator(),generator()
	var worker = createWorker(0)

	for {
		select {
		case n := <- c1:
			worker <- n
		case n := <- c2:
			worker <- n
		}
	}
}
/**
	当有数据产生的时候，在处理
 */
func selectCaseV3()  {
	var c1,c2 = generator(),generator()
	var worker = createWorker(0)
	var values []int //将生成的数据保存到这个数组中，以免因为程序处理过快，漏掉部分数据
	for {
		var activeWorker chan <- int
		var activeValue int //当前要被处理的数据
		//当接收到了数据后，开始执行任务
		if len(values) > 0{
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <- c1:
			values = append(values,n)
		case n := <- c2:
			values = append(values,n)
		case activeWorker <- activeValue:  //当要被处理的数据有值的时候，将里面的数据提出一个
			values = values[1:]
		}
	}
}