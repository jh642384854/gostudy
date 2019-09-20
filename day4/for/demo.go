package main

import "fmt"

//基于计数器的for循环
func Test1()  {
	for i := 0; i< 10 ;i++  {
		fmt.Println("i = ",i)
	}
}

//基于条件的for循环(和其他编程语言的while类似)
func Test2()  {
	num := 10
	for num > 0  {
		num = num -1
		fmt.Println("num = ",num)
	}
}

//无限循环
func Test3()  {
	var (
		num1 = 1
		num2 = 1
		num3 = 1
	)
	infinite_loop1(num1)
	infinite_loop2(num2)
	infinite_loop3(num3)
}
//无限循环形式1
func infinite_loop1(num1 int)  {
	for ; ;  {
		fmt.Println("num1 = ",num1)
		num1 = num1 +1
		if num1 > 10{
			return
		}
	}
}
//无限循环形式2
func infinite_loop2(num2 int)  {
	for i:=0;;i++{
		fmt.Println("num2 = ",num2)
		num2 = num2 +1
		if num2 > 10{
			return
		}
	}
}
//无限循环形式3
func infinite_loop3(num3 int)  {
	for{
		fmt.Println("num3 = ",num3)
		num3 = num3 +1
		if num3 > 10{
			return
		}
	}
}

//for  range结构
func Test4()  {
	for i,j,s := 0,5,"a"; i<3 && j<100 && s!= "aaaaa" ; i,j,s = i+1,j+1,s+"a"  {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}

func main() {
	//Test1()
	//Test2()
	//Test3()
	Test4()
}
