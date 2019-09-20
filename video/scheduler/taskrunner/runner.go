package taskrunner

type Runner struct {
	Controller controllerChan   //定义当前控制类别
	Error controllerChan        //记录错误信息
	Data dataChan               //真正需要处理的数据
	Size int                    //同时可以被处理的数据量
	LongLived bool              //
	Dispatch fn                 //分发器处理函数
	Execute fn                  //处理数据函数
}

func NewRunner(size int,longlived bool,d fn,e fn) *Runner {
	return &Runner{
		Controller:make(chan string,1),
		Error:make(chan string,1),
		Data:make(chan interface{},size),
		Size:size,
		LongLived:longlived,
		Dispatch:d,
		Execute:e,
	}
}

func (r *Runner) startDispacth()  {
	defer func() {
		if !r.LongLived{
			close(r.Controller)
			close(r.Error)
			close(r.Data)
		}
	}()
	//使用常驻任务来执行
	for  {
		select {
		case controllerType := <- r.Controller:
			if controllerType == READY_TO_DISPATCH{
				err := r.Dispatch(r.Data)
				if err != nil{
					r.Error <- CLOSE
				}else{
					r.Controller <- READY_TO_EXECUTE
				}
			}
			if controllerType == READY_TO_EXECUTE{
				err := r.Execute(r.Data)
				if err != nil{
					r.Error <- CLOSE
				}else{
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case err := <- r.Error:
			if err == CLOSE{
				return
			}
		default:

		}
	}
}

func (r *Runner) StartAll()  {
	//首先就开启一个分发任务
	r.Controller <- READY_TO_DISPATCH
	r.startDispacth()
}