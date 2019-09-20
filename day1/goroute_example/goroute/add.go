package goroute

//这里用到来管道
func Add(a int, b int, c chan int) {
	sum := a + b
	c <- sum
}
