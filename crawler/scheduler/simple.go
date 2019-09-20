package scheduler

import "dev/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan  engine.Request
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(c chan engine.Request) {
	s.WorkerChan = c
}

//直接发送请求到worker chan中
func (s *SimpleScheduler) Submit(r engine.Request) {
	//这里避免出现循环等待，开启一个goroutine来处理
	go func() {
		s.WorkerChan <- r
	}()
}
