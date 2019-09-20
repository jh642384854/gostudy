package main

import (
	"fmt"
	"runtime"
)

/**
	需求：
	开启4个协程(其实是根据电脑的CPU的核心数来决定)来统计 1-200000 的数字中，哪些是素数？具体实现过程如下：

	①、专门定义一个方法，用来将200000个数字写到numChan的channel中
	②、专门定义一个方法，用来从numChan的channel通道中取出数据，并判断这个数据是否是素数
 */

func WriteNumData(numChan chan int)  {
	for i:= 0;i<=200000 ; i++ {
		numChan <- i
	}
	close(numChan)
}

func ReadNumData(numChan chan int,resChan chan int,exitChan chan bool)  {
	var flag bool
	for {
		num,ok := <- numChan
		if !ok{
			break
		}
		flag = true
		//判断是否为素数
		for i:=2;i<num ;i++  {
			if num %i == 0{
				flag = false
				break
			}
		}
		if flag{
			resChan <- num
		}
	}
	exitChan <- true
}

func main() {

	cpuNumbers := runtime.NumCPU()

	fmt.Println("当前系统拥有",cpuNumbers,"个CPU")

	numChan := make(chan int,1000)

	resChan := make(chan int,1000)

	exitChan := make(chan bool,cpuNumbers)

	go WriteNumData(numChan)
	//开启CPU的核心数个协程来处理数据
	for i:=0;i<cpuNumbers ;i++  {
		go ReadNumData(numChan,resChan,exitChan)
	}

	go func() {
		for i:=0;i<cpuNumbers ;i++  {
			<- exitChan
		}
		//关闭统计素数结果的channel
		close(resChan)
	}()

	for {
		_,ok := <- resChan
		if !ok{
			break
		}
		//fmt.Println("素数结果为：",v)
	}

	fmt.Println("主线程运行退出")

}
