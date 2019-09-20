package main

import "fmt"

/**
	map的创建和初始化
 */

//map的几种创建方式
func Test1()  {
	var map1 map[int]string
	map2 := map[int]string{}
	map3 := make(map[int]string)
	map4 := make(map[int]string,10)
	fmt.Println(map1,map2,map3,map4,len(map4))
}

//map的初始化方式
func Test2()  {
	var map1 map[int]string = map[int]string{1:"hello",2:"world"}
	map2 := map[int]string{1:"demo1",2:"demo2"}


	map3 := map[int]string{}
	map3[2] = "test1"
	map3[3] = "test2"

	map4 := make(map[int]string)
	map4[1] = "map4 value 1"
	map4[2] = "map4 value 2"


	fmt.Println(map1,map2,map3,map4)
}

func main() {
	Test1()
	Test2()
}
