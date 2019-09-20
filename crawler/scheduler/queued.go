package scheduler

import "dev/crawler/engine"

type QueuedScheduler struct {
	RequestChan chan engine.Request      //请求队列
	WorkerChan  chan chan engine.Request //worker 队列
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.RequestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.WorkerChan <- w
}

func (s *QueuedScheduler) ConfigMasterWorkerChan(c chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) Run() {
	s.WorkerChan = make(chan chan engine.Request)
	s.RequestChan = make(chan engine.Request)
	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorkder chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0{
				activeWorkder = workerQueue[0]
				activeRequest = requestQueue[0]
			}
 			select {
			case r := <-s.RequestChan:
				requestQueue = append(requestQueue,r)
			case w := <-s.WorkerChan:
				workerQueue = append(workerQueue,w)
			//从队列拿掉数据
			case activeWorkder <- activeRequest:
				workerQueue = workerQueue[1:]
				requestQueue = requestQueue[1:]
			}
		}
	}()
}
