package main

import (
	"fmt"
	"time"
)

/**
	本示例主要演示关闭通道，或往通道写入特定值来结束某个任务
 */

func main() {
	TestCancel()
}

//向这个chan中写入一个空的结构体对象，因为这里没有太多实际数据要写入，如果实际有什么业务需求，可以单独来实现。
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

//关闭通道
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

//判断任务是否被取消，就是通过读取cancelChan这个channel的值
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func TestCancel() {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				//这里循环执行，可以假定为要长时间执行的任务。下面有一个检查任务是否被取消，如果检测到任务被取消，就需要终止该任务的执行
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Done")
		}(i, cancelChan)
	}
	cancel_1(cancelChan)  //使用下面的方法也可以
	//cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}
