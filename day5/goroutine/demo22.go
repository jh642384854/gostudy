package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg用来等待程序结束
var wg6 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建一个无缓冲通道
	court := make(chan int)

	//计数加2，表示要等待2个goroutine
	wg6.Add(2)
	go player("zhagnsan",court)
	go player("lisi",court)
	//发球
	court <- 1
	//等待游戏结束
	wg6.Wait()
}

// 模拟一个选手打球
func player(name string,court chan int)  {
	defer wg6.Done()

	for  {
		//等待球被击打过来
		ball,ok := <- court
		if !ok {
			//如果通道被关闭，我们就赢了
			fmt.Printf("Player %s won \n",name)
			return
		}
		//选随机数，然后用这个数据来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0{
			fmt.Printf("Player %s Missed \n",name)
			//关闭通道，表示我们输了
			close(court)
			return
		}
		fmt.Printf("Player %s Hits %d \n",name,n)
		ball ++
		court <- ball
	}
}