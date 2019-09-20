package main

import (
	"context"
	"time"
	"fmt"
)

/**
	context控制多个goroutine

	示例中启动了3个监控goroutine进行不断的监控，每一个都使用了Context进行跟踪，当我们使用cancel函数通知取消时，这3个goroutine都会被结束。这就是Context的控制能力，它就像一个控制器一样，按下开关后，所有基于这个Context或者衍生的子Context都会收到通知，这时就可以进行清理操作了，最终释放goroutine，这就优雅的解决了goroutine启动后不可控的问题。
 */

func main() {
	ctx,cancel := context.WithCancel(context.Background())

	go watch(ctx,"【监控1】")
	go watch(ctx,"【监控2】")
	go watch(ctx,"【监控3】")

	time.Sleep(5*time.Second)
	fmt.Println("可以了，通知监控停止")
	defer cancel()
	time.Sleep(5*time.Second)
}

func watch(ctx context.Context,name string)  {
	for  {
		select {
		case <- ctx.Done():
			fmt.Println("监控退出，停止了...",ctx.Err())
			return
		default:
			fmt.Println("go routine持续监控中...")
			time.Sleep(2*time.Second)
		}
	}
}