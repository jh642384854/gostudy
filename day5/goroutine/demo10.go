package main

import "fmt"

/**
	串行channel
	需求说明：
	定义两个channel(nums和squares)，其中nums中用于存放一组数据。而squares里面存放的是nums中每个数据的平方
	然后在主线程中打印出squares数据
 */

 /**
 	像这样的串联Channels的管道（Pipelines）可以用在需要长时间运行的服务中，
 	每个长时间运行的goroutine可能会包含一个死循环，在不同goroutine的死循环内部使用串联的Channels来通信。
 	但是，如果我们希望通过Channels只发送有限的数列该如何处理呢？
  */
func errDemo()  {
	nums := make(chan int)

	//生成100个数字，将这些数字放入到nums中
	go func() {
		for i := 1;i<= 100 ; i++ {
			nums <- i
		}
		//这里没有关闭channel，在读取数据的时候，会出问题
	}()

	//读取nums中的值，然后做平方运算
	squares := make(chan int)
	go func() {
		//这种方式读取数据会出问题，因为没有关闭channel
		for  {
			x  := <- nums
			squares <- x*x
		}
		//这里没有关闭channel，在读取数据的时候，会出问题
	}()

	//在主线程打印squares的值
	for  {
		fmt.Println(<- squares)
	}
}

func errDemo2()  {
	nums := make(chan int)

	//生成100个数字，将这些数字放入到nums中
	go func() {
		for i := 1;i<= 100 ; i++ {
			nums <- i
		}
	}()

	//读取nums中的值，然后做平方运算
	squares := make(chan int)
	go func() {
		//通过for循环的方式，来获取channel的元素
		for {
			v,ok := <- nums
			if !ok{
				break
			}
			squares <- v*v
		}
		close(squares)//关闭channel
	}()

	//在主线程打印squares的值
	for  {
		fmt.Println(<- squares)
	}
}
/**
	其实你并不需要关闭每一个channel。只要当需要告诉接收者goroutine，所有的数据已经全部发送时才需要关闭channel。
	不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾自动回收器回收。
	（不要将关闭一个打开文件的操作和关闭一个channel操作混淆。对于每个打开的文件，都需要在不使用的使用调用对应的Close方法来关闭文件。）
	试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常。关闭一个channels还会触发一个广播机制，
 */
func SuccDemo()  {
	nums := make(chan int)

	//生成100个数字，将这些数字放入到nums中
	go func() {
		for i := 1;i<= 100 ; i++ {
			nums <- i
		}
		close(nums)                                  //重点1：需要关闭管道
	}()

	//读取nums中的值，然后做平方运算
	squares := make(chan int)
	go func() {
		//通过for range的方式来遍历channel中的值       重点2：使用for range的方式来遍历channel
		//这里需要注意一点，一般使用for range方式会得到两个值，一个是索引，一个是value，而channel的值是没有索引的，只有value，所以在进行遍历的时候，也就只有value值返回。这点需要注意
		for value := range nums {
			squares <- value*value
		}
		close(squares)                                 //重点3：需要关闭管道
	}()

	//在主线程打印squares的值                          //重点4：使用for range的方式来遍历channel
	for value := range squares {
		fmt.Println(value)
	}
}


/**
	调用counter(naturals)将导致将 chan int 类型的naturals隐式地转换为 chan<- int 类型只发送型的channel。调用printer(squares)也会导致相似的隐式转换，
	这一次是转换为 <-chanint 类型只接收型的channel。任何双向channel向单向channel变量的赋值操作都将导致该隐式转换。
	这里并没有反向转换的语法：也就是不能将一个类似 chan<- int 类型的单向型的channel转换为 chan int 类型的双向型的channel
 */
func main() {
	nums := make(chan int)
	squares := make(chan int)
	go counter(nums)
	go squaresf(nums,squares)
	printer(squares)
}

//专门用来生成数据的方法，并将数据写入到channel中。
//注意这里形参的写法numser chan <- int，表示这个channel只能是一个写入的channel
func counter(nums chan <- int) {
	for i := 1;i<=100 ;i++  {
		nums <- i
	}
	close(nums)
}
//将生成的数据逐个进行平方运算
//out <- chan int这个就表示，out这个channel是一个只读channel。
//in chan <- int这个就表示，in这个channel是一个只写channel。
func squaresf(out <- chan int,in chan <- int)  {
	for value := range out {
		in <- value*value
	}
	close(in)
}
//打印数据
func printer(squares <- chan int)  {
	for value := range squares {
		fmt.Println(value)
	}
}