package main

import (
	"fmt"
	"sync"
)

/**
	channel的关闭和广播
 */

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataConsumer(ch, &wg)

	//可以创建多个consumer来消费同一个producer
	wg.Add(1)
	dataConsumer(ch, &wg)

	wg.Wait()
}

//数据生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//完成数据写入后，就关闭通道。注意，对于关闭的通道是无法继续往通道里面写数据的
		close(ch)
		wg.Done()
	}()
}

//数据消费者
func dataConsumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			//从通道里面获取数据，根据获取的值来判断是否正常获取到值
			if data, err := <-ch; err {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}
