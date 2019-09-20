package main

import "fmt"

/**
	只读/只写 channel

	默认情况下，创建的channel是双向的，即可读、也可写
	如果定义的channel是只读或是只写的，就需要在声明的时候指定。

	声明方式的区别：
	只写channel：chan<- int
	只读channel：<-chan int

	特别需要说明一个事情：我们定义的channel是只读/只写，只是定义这个channel的属性。就是我们可以创建一个channel，然后把这个channel当作形参传值给一个只读/只写的参数
 */

//只写channel
func WriteOnlyChannel()  {

	woChan := make(chan<- int,4)

	woChan <- 1
	woChan <- 2
	woChan <- 3
	woChan <- 4
	close(woChan)

	//num := <- woChan  //创建的channel是只写模式的话，在读取的话就会报错
}

//只读channel
func ReadOnlyChannel()  {
	roChan := make(<-chan int,4)

	nums := <- roChan   //这样从channel中读取值是正常的
	fmt.Println(nums)

	//roChan <- 1      //这样往channel中写数据就会出问题
}

/**
	我们定义了一个可读可写的channel，但是我们把这个channel传递一个函数里面，而这个函数里面的两个参数一个是只写的channel，一个是只读的channel。
	所以，我们所说的可读/可写的channel，其实是一种特殊的属性，而并不是数据类型本身。
 */
func TestRWChan(woChan chan<- int,roChan <-chan int)  {
	//写数据
	woChan <- 11
	woChan <- 22
	woChan <- 33
	woChan <- 44
	woChan <- 55

	//读数据
	nums1 := <- roChan
	fmt.Println(nums1)
}

func main() {

	numChan := make(chan int,5)
	/*numChan <- 4
	numChan <- 5
	numChan <- 6
	numChan <- 7
	numChan <- 8*/

	TestRWChan(numChan,numChan)


	//下面的示例，展示了只读/只写channel函数的使用
	var ronly = readOnly()
	var ronlys  chan<- int
	ronlys = ronly
	ronlys <- 15
}

//定义的只读channel的函数
func readOnly() chan <- int {
	c := make(chan int)
	go writeOnly(c)
	return c
}

//定义的只写channel函数
func writeOnly(c <-chan int)  {
	for value := range c {
		fmt.Println(value)
	}
}