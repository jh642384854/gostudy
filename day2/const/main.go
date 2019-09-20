package main

import "fmt"

func test() {
	// 常量只能通过const关键字来声明定义。常量可以定义为数值、布尔值或字符串等类型
	const PI float32 = 3.1415926
	const i = 1000
	const MaxThread = 10
	const prefix = "string"
	fmt.Println(PI,i,MaxThread,prefix)
}

func main() {
	test()
}