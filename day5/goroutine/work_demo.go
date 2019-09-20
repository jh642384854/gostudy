package main

import (
	"dev/day5/work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"zhangsan","lisi","wangwu","zhaoliu","wangmazi",
}

type namePrinter struct {
	name string
}

//Task实现Worker接口
func (n *namePrinter) Task()  {
	log.Println(n.name)
	time.Sleep(time.Second)
}

func main() {
	//使用两个goroutine来创建工作池
	p := work.New(2)

	var wg9 sync.WaitGroup
	wg9.Add(10*len(names))

	for i := 0;i<10;i++{
		//迭代names切片
		for _, name := range names {
			//创建一个namePrinter并提供指定的名字
			np := namePrinter{
				name:name,
			}
			go func() {
				p.Run(&np)
				wg9.Done()
			}()
		}
	}
	wg9.Wait()
	//让工作池停止工作，等待所有现有的工作完成
	p.Shutdown()
}