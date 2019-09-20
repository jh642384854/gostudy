package main

import (
	"context"
	"fmt"
	"time"
)

/**
	本示例主要演示context的使用

	根Context：通过context.Background()来创建
	子Context：通过context.WithCancel(parentContext)来创建
		示例：ctx,cancel := context.WithCancel(context.Background())
	当前context被取消时候，基于它的子context或孙等等其他context都会被取消。
	接收取消的通知： <- ctx.Done()
 */

func main() {
	TestCancel()
}

//判断任务是否被取消，就是通过读取cancelChan这个channel的值
func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, context context.Context) {
			for {
				//检查任务是否被取消
				if isCancelled(context) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Done")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
