package taskrunner

import (
	"log"
	"testing"
	"time"
)

func TestNewRunner(t *testing.T) {
	//分发数据
	dispatch := func(dc dataChan) error {
		for i := 0;i<30 ;i++  {
			dc <- i
			log.Printf("Dispatch send :%v",i)
		}
		return  nil
	}
	//处理数据
	execute := func(dc dataChan) error {
		forloop:
			for  {
				select {
				case data := <- dc:
					log.Printf("Executor receive :%v",data)
				default:
					break forloop
				}
			}
		return nil
		//return errors.New("Executor")
	}
	runner := NewRunner(30,false,dispatch,execute)
	go runner.StartAll()
	time.Sleep(1* time.Second)
}

func TestDemo(t *testing.T)  {
	go func() {
		for {
			for i := 0; i<= 30 ;i++  {
				log.Printf("get i :%v",i)
			}
		}
	}()
	time.Sleep(3*time.Second)
}