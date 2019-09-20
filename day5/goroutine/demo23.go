package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	GoroutinesNumbsers = 4 //要使用的goroutine数量
	TaskLoad = 10  // 要处理的工作数量
)

// wg用来等待程序结束
var wg7 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//创建一个带有缓冲的通道来管理工作
	tasks := make(chan string,TaskLoad)
	//启动goroutine来处理工作
	wg7.Add(GoroutinesNumbsers)
	for gr := 1;gr <= GoroutinesNumbsers ; gr++  {
		go worker(tasks,gr)
	}
	//增加一组要完成的工作。向通道里写数据
	for post := 1; post <= TaskLoad ;post++  {
		tasks <- fmt.Sprintf("Task : %d",post)
	}
	//当所有工作都处理完成时候关闭通道，以便所有的goroutine退出
	close(tasks)
	//等待所有工作完成
	wg7.Wait()
}
func worker(tasks chan string,worker int)  {
	defer wg7.Done()
	for  {
		task ,ok := <- tasks
		if !ok {
			//这意味着通道已经空了，并且已经关闭
			fmt.Printf("Worker :%d Shutting Down \n",worker)
			return
		}
		//显示我们开始工作了
		fmt.Printf("Workder : %d Started %s\n",worker,task)
		//随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep)*time.Millisecond)
		//显示我们完成了工作
		fmt.Printf("Workder : %d Completed %s",worker,task)
	}
}



























