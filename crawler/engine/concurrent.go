package engine

/**
	并发处理方式
 */

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	//interface里面的函数参数不需要命名
	Submit(Request)                      //将获取的url发送到worker chan中
	ConfigMasterWorkerChan(chan Request) //定义那个是worker chan

	WorkerReady(w chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult) //处理结果的channel

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	//接收处理的结果
	for {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("Got Item:%v \n",item)
			//将要保存的信息存放到ItemChan 这个channel中
			e.ItemChan <- item
		}
		//将重新获取的地址添加到队列里面
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan ParseResult, scheduler Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			scheduler.WorkerReady(in)
			request := <-in
			result, e := worker(request)
			if e != nil {
				continue
			}
			out <- result
		}
	}()
}

/*
SimpleScheduler的调用方式
func (e *ConcurrentEngine) Run(seeds ...Request)  {
	in := make(chan Request)     //请求URL地址的channel
	out := make(chan ParseResult)//处理结果的channel

	e.Scheduler.ConfigMasterWorkerChan(in)

	for i := 0; i<e.WorkerCount ;i++  {
		createWorker(in,out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	//接收处理的结果
	for  {
		result := <- out
		for _, item := range result.Items {
			log.Printf("Got Item:%v \n",item)
		}
		//将重新获取的地址添加到队列里面
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult)  {
	go func() {
		for  {
			request := <- in
			result, e := worker(request)
			if e != nil{
				continue
			}
			out <- result
		}
	}()
}
*/
