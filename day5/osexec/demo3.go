package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

/**
	这个示例展示了如何杀死一个进程(就是结束掉一个正在执行的命令)

	需求如下：
	执行一个CMD,让它在一个协程里面去执行，让它执行2秒：sleep 2；然后在输出hello
	当过了1秒之后，我们就杀死这个CMD，最后的结果就是不会输出hello
 */

type result struct {
	err    error
	output []byte
}

func main() {
	var (
		//定义一个cmd对象变量
		cmd *exec.Cmd
		//定义一个ctx上下文对象变量
		ctx context.Context
		//定义一个context.CancelFunc变量
		cancelFun context.CancelFunc
		//定义一个记录下面协程执行结果的变量
		resultChan chan *result
		//定义一个从协程里面获取结果的变量
		res *result
	)
	resultChan = make(chan *result)
	//创建一个上下文对象
	ctx, cancelFun = context.WithCancel(context.TODO())

	go func() {
		var (
			//定义cmd.CombinedOutput()输出结果变量
			output []byte
			//定义在执行cmd.CombinedOutput()出错后的错误信息
			err error
		)
		//创建一个cmd对象
		cmd = exec.CommandContext(ctx, "C:\\cygwin64\\bin\\bash.exe", "-c", "sleep 2;ls -l")
		//执行任务，并捕获结果
		output, err = cmd.CombinedOutput()
		//将结果输出到channel，在main协程里面来获取数据
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()
	//模拟程序执行1秒
	time.Sleep(1 * time.Second)
	//取消上下文
	cancelFun()
	//得到结果
	res = <-resultChan
	fmt.Println(res.err, string(res.output))
}
