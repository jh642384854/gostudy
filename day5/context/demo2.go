package main

import (
	"context"
	"fmt"
	"time"
)

/**
	上面说的这种场景是存在的，比如一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，这些goroutine又可能会开启其他的goroutine。所以我们需要一种可以跟踪goroutine的方案，才可以达到控制他们的目的，这就是Go语言为我们提供的Context，称之为上下文非常贴切，它就是goroutine的上下文。
	重写比较简单，就是把原来的chan stop 换成Context，使用Context跟踪goroutine，以便进行控制，比如结束等。
	context.Background() 返回一个空的Context，这个空的Context一般用于整个Context树的根节点。然后我们使用context.WithCancel(parent)函数，创建一个可取消的子Context，然后当作参数传给goroutine使用，这样就可以使用这个子Context跟踪这个goroutine。
	在goroutine中，使用select调用<-ctx.Done()判断是否要结束，如果接受到值的话，就可以返回结束goroutine了；如果接收不到，就会继续进行监控。
	那么是如何发送结束指令的呢？这就是示例中的cancel函数啦，它是我们调用context.WithCancel(parent)函数生成子Context的时候返回的，第二个返回值就是这个取消函数，它是CancelFunc类型的。我们调用它就可以发出取消指令，然后我们的监控goroutine就会收到信号，就会返回结束。
 */

func main() {

	ctx,cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for  {
			select {
				case <- ctx.Done():
					fmt.Println("监控退出，停止了...")
					return
			default:
				fmt.Println("go routine持续监控中...")
				time.Sleep(2*time.Second)
			}
		}
	}(ctx)

	time.Sleep(5*time.Second)
	fmt.Println("可以了，通知监控停止")
	defer cancel()
	time.Sleep(5*time.Second)
}