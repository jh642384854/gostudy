package main

import (
	"fmt"
	"time"
)

/**
	所有任务都完成后，并合并处理
	①、使用sync.WaitGroup的方式来实现
	②、使用csp的机制来实现(即channel)
 */

func main() {
	fmt.Println(AllResponse())
}



func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numOfRunner := 10

	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(num int) {
			res := runTask(num)
			ch <- res
		}(i)
	}
	//等待上面任务执行完毕后进行合并处理
	finalRes := ""
	for i:=0;i<numOfRunner ;i++  {
		finalRes += <-ch + "\n"
	}
	return finalRes
}