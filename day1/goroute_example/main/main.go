package main

import (
	"dev/day1/goroute_example/goroute"
	"fmt"
)

//这个函数会优先main()函数执行
func init() {

}

func main() {
	//声明一个管道，然后把计算结果放在管道里面
	var pipe chan int
	pipe = make(chan int, 1)
	go goroute.Add(100, 200, pipe)
	//从管道取出数据
	sum := <-pipe
	fmt.Println("sum = ", sum)
}
