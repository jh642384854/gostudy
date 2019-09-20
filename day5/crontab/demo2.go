package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

/**
	定义一个多任务执行的
 */
//定义一个任务结构体对象
type crontabJob struct {
	expr     *cronexpr.Expression
	nexttime time.Time
}

func main() {
	var (
		expr        *cronexpr.Expression
		nextTime    time.Time
		nowTime     time.Time
		crontabJobs map[string]*crontabJob
	)
	nowTime = time.Now()
	//map类型需要先make，才能往里面添加值
	crontabJobs = make(map[string]*crontabJob)
	//创建任务1
	expr = cronexpr.MustParse("*/5 * * * * * *")
	nextTime = expr.Next(nowTime)
	crontabJobs["job1"] = &crontabJob{
		expr:     expr,
		nexttime: nextTime,
	}
	//创建任务2
	expr = cronexpr.MustParse("*/5 * * * * * *")
	nextTime = expr.Next(nowTime)
	crontabJobs["job2"] = &crontabJob{
		expr:     expr,
		nexttime: nextTime,
	}

	//开启协程来处理这些任务
	go func() {
		var (
			job     *crontabJob
			jobName string
			now     time.Time
		)
		for {
			now = time.Now()
			for jobName, job = range crontabJobs {
				//如果到了定时执行时间，就需要执行相应的任务了
				if job.nexttime.Before(now) || job.nexttime.Equal(now) {
					//启动一个协程，去执行这个任务
					go func(jobName string) {
						fmt.Println(jobName)
					}(jobName)
					//计算下一次任务执行时间
					job.nexttime = job.expr.Next(now)
					fmt.Println("下次执行时间是：", job.nexttime)
				}
			}
			//睡眠100毫秒
			select {
			//这里就等同于time.sleep(100 * time.Millisecond)
			case <-time.NewTimer(100 * time.Millisecond).C:
			}
		}
	}()

	time.Sleep(100 * time.Second)
}
