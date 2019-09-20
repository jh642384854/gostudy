package main

import (
	"dev/day5/runner"
	"log"
	"os"
	"time"
)

//timeout规定了必须要在多少秒内处理完成
const timeout = 3 * time.Second

func main() {

	log.Println("Starting Work.")

	//为本次执行分配超时时间
	r := runner.New(timeout)

	//添加要执行的任务
	r.Add(createTask(), createTask(), createTask())

	//执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeOut:
			log.Println("Termination due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Termination due to interrupt.")
			os.Exit(2)
		}
	}
	log.Println("Process ended")
}

func createTask() func(int) {
	return func(i int) {
		log.Printf("Processor - Task #%d.", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}
