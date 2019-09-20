package main

import "fmt"

//判断map中是否存在指定的key
func TestFun1()  {
	map1 := make(map[int]string)
	map1[1] = "demo1"
	map1[2] = "demo2"

	var val,isPresent = map1[2]

	if isPresent{
		fmt.Println("map1[2]:",val)
	}

	//如果只想判断指定key是否存在，可以用下面的写法
	if _,isPresent := map1[3];isPresent{
		fmt.Println(map1[3])
	}
}

//删除map中的key
func TestFun2()  {
	map1 := make(map[int]string)
	map1[1] = "demo1"
	map1[2] = "demo2"

	delete(map1,3)//删除指定的key即便不存在，也不会报错
	delete(map1,2)
	fmt.Println(map1)
}

func main() {
	TestFun1()
	TestFun2()
}