package runner

import (
	"github.com/pkg/errors"
	"os"
	"os/signal"
	"time"
)
/**
	runner 包用于展示如何使用通道来监视程序的执行时间，如果程序运行时间太长，也可以用 runner 包来终止程序。
	当开发需要调度后台处理任务的程序的时候，这种模式会很有用。这个程序可能会作为 cron 作业执行，或者在基于定时任务的云环境（如 iron.io）里执行
 */

//runner包管理处理任务的运行和生命周期

// Runner 在给定的超时时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	//interrupt通道报告从操作系统发送的信号
	interrupt chan os.Signal

	//complete通道报告任务已经完成
	complete chan error

	// timeout报告处理任务已经超时
	timeout <-chan time.Time

	//tasks持有一组以索引顺序依次执行的函数(任务)
	tasks []func(int)
}

// ErrTimeOut会在任务执行超时时返回
var ErrTimeOut = errors.New("received timeout")

// ErrInterrupt会在接收到操作系统的事件返回
var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//添加一个任务到Runner上，这个任务是接收一个int类型ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// 执行所有的任务，并监视通道事件
func (r *Runner) Start() error {
	//我们希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	//用不同的goroutine来执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

//run执行每一个已经注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//检测到系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

//验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	//当中断事件被触发时发出的信号
	case <-r.interrupt:
		//停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
	default:
		//继续正常运行
		return false
	}
}
