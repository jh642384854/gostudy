package main

import (
	"fmt"
	"time"
)

/**
	channel 的select用法

	使用 select case 可以解决从管道取数据的阻塞问题

	前面我们了解到的channel，要遍历channel的值，必须是要执行close()之后，才可以，但是，有些时候，我们并不知道什么时候执行close()，这样有没有别的好办法在没有close()的情况下来遍历channel呢？
 */

func main() {

	intChan := make(chan int,5)
	strChan := make(chan string,5)

	//向channel中添加数据，但是我们并没有执行close()
	for i:=0;i<5 ;i++  {
		intChan <- i*2
		strChan <- fmt.Sprintf("value:%v",i)
	}

	for {
		select {
			//注意: 这里，如果 intChan 一直没有关闭，不会一直阻塞而 deadlock ,会自动到下一个 case 匹配
			case v:= <- intChan:
				fmt.Println("intChan value:",v)
				//time.Sleep(time.Second)
			case v:= <- strChan:
				fmt.Println("strChan value:",v)
				time.Sleep(time.Second)
			default:
				fmt.Println("数据都取完毕了")
				return //这里的return非常重要，用来退出循环
		}
	}
}

