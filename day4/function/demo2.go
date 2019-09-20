package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func f() (ret int)  {
	defer func() {
		ret ++
	}()
	return 1
}
//函数的返回值也是一个匿名函数(闭包)
func adder(a int) func(b int) int  {
	return func(b int) int {
		return a+b
	}
}

func add2() func(b int) int  {
	return func(b int) int {
		return  b+2
	}
}

func add3() func(int) int  {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}

func MakeAddSuffix(suffix string) func(file string) string  {
	return func(file string) string {
		//debug()
		if !strings.HasSuffix(file,suffix){
			//debug()
			return file + suffix
		}
		return file
	}
}
//代码调试
func debug()  {
	_,file,line,_ := runtime.Caller(1)
	log.Printf("%s:%d",file,line)
}

func demo(n int) (res int)  {
	fmt.Println("demo():n is",n)
	if n <=1 {
		res = 1
	}else{
		res = demo(n-1)+demo(n-2)
	}
	return
}
func main() {
	fmt.Println(f())
	num1 := adder(10)
	fmt.Println(num1(5))
	num2 := add2()
	fmt.Println(num2(5))

	f := add3()
	fmt.Println(f(1))
	fmt.Println(f(10))
	fmt.Println(f(20))

	jpg := MakeAddSuffix(".jpg")
	png := MakeAddSuffix(".png")
	fmt.Println(jpg("a.jpg"))
	fmt.Println(png("demo"))
	fmt.Println("demo():")
	fmt.Println(demo(7))
}