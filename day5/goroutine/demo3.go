package main

import "fmt"

/**
	channel入门示例
	channel：数据是遵循先进先出原则(就是一种队列数据类型)。就是先添加进去的数据，在取出来的时候，也会有优先取出来。
	刚好和栈相反
 */

func main() {
	//创建一个可以存放3个int类型的管道
	var intChan chan int

	intChan = make(chan int,3)
	//看看inchan是啥？
	fmt.Printf("intChan的值是：%v，intChan的地址是：%v\n",intChan,&intChan)

	//向管道写入数据
	intChan <- 10
	//创建一个变量，然后将这个变量写入管道中
	num := 123
	intChan <- num

	intChan <- 555

	//这里已经给inChan管道写入了3个元素，这样就无法在写入数据了。这个不像slice会动态的增加元素长度
	//但是我们可以做到边写、边读取，只要写入的数据长度不超过创建的chan长度，这样是可以的

	//我们在来看看管道的长度和容量
	fmt.Printf("管道的长度：%v,管道的容量:%v \n",len(intChan),cap(intChan))

	//读取管道数据

	num1 := <- intChan
	num2 := <- intChan
	num3 := <- intChan

	fmt.Printf("num1: %v,num2:%v,num3:%v \n",num1,num2,num3)

}