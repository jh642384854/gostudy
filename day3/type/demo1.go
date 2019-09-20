package main

import "fmt"

type DiyFun func(i,j int) int

type Integer int

type numbsers []int

func main() {
	diyfun := func(a,b int) int {
		return a + b
	}
	result := diyfun(2,3)
	fmt.Println(result)

	fmt.Println(Integer(20))

	fmt.Println(numbsers([]int{1,2,3,4}))

	fmt.Println(DiyFun(diyfun)) //当定义的类型是一个func的时候，一定要注意这种调用写法。
}
