package taskrunner

import "time"

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration,r *Runner) *Worker {
	return &Worker{
		ticker:time.NewTicker(interval),
		runner:r,
	}
}

func (w *Worker) startWorker()  {
	for  {
		select {
		case  <- w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start()  {
	runner := NewRunner(10,true,VideoClearDispatcher,VideoClearExecutor)
	workder := NewWorker(5,runner)
	workder.startWorker()
}