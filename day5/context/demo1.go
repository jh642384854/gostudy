package main

import (
	"fmt"
	"time"
)

/**
	context的引入：

	我们都知道一个goroutine启动后，我们是无法控制他的，大部分情况是等待它自己结束，那么如果这个goroutine是一个不会自己结束的后台goroutine呢？比如监控等，会一直运行的。
	这种情况化，一直傻瓜式的办法是全局变量，其他地方通过修改这个变量完成结束通知，然后后台goroutine不停的检查这个变量，如果发现被通知关闭了，就自我结束。
	这种方式也可以，但是首先我们要保证这个变量在多线程下的安全，基于此，有一种更好的方式：chan + select 。

	例子中我们定义一个stop的chan，通知他结束后台goroutine。实现也非常简单，在后台goroutine中，使用select判断stop是否可以接收到值，如果可以接收到，就表示可以退出停止了；如果没有接收到，就会执行default里的监控逻辑，继续监控，只到收到stop的通知。
	有了以上的逻辑，我们就可以在其他goroutine种，给stop chan发送值了，例子中是在main goroutine中发送的，控制让这个监控的goroutine结束。
	发送了stop<- true结束的指令后，我这里使用time.Sleep(5 * time.Second)故意停顿5秒来检测我们结束监控goroutine是否成功。如果成功的话，不会再有goroutine监控中...的输出了；如果没有成功，监控goroutine就会继续打印goroutine监控中...输出。
	这种chan+select的方式，是比较优雅的结束一个goroutine的方式，不过这种方式也有局限性，如果有很多goroutine都需要控制结束怎么办呢？如果这些goroutine又衍生了其他更多的goroutine怎么办呢？如果一层层的无穷尽的goroutine呢？这就非常复杂了，即使我们定义很多chan也很难解决这个问题，因为goroutine的关系链就导致了这种场景非常复杂。

	https://www.flysnow.org/2017/05/12/go-in-action-go-context.html

*/

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("go routine持续监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	time.Sleep(5 * time.Second)
}
