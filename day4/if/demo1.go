package main

import "fmt"

func Test1() {
	num := 11
	if num > 10 {
		fmt.Println("num gt 10", num);
	} else {
		fmt.Println("num lt 10");
	}
}
// bool值直接判断
func Test2()  {
	bool1 := false
	if bool1 {
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}
}
//多if else if分支
func Test3()  {
	method := 4
	if method > 3 && method < 6 {
		fmt.Println("Spring")
	}else if method >6 && method <9{
		fmt.Println("summer")
	}else if method >6 && method <9{
		fmt.Println("autumn")
	}else{
		fmt.Println("winter")
	}
}
//先赋值后做判断
func Test4()  {
	if num := 51; num > 41{
		fmt.Println("num gt 41")
	}else{
		fmt.Println("num lt 41")
	}
}
// 先经过函数执行，在做判断
func Test5()  {
	number := 15
	if process(number){
		fmt.Println("num gt 10")
	}else{
		fmt.Println("num lt 10")
	}
}
//传递的参数必须要指定数据类型，如果有返回值的情况下，还需要指定返回值类型
func process(num int) bool  {
	if num > 10{
		return true
	}else{
		return false
	}
}
// 赋值通过函数执行，然后在做判断
func Test6()  {
	if max := process2(); max > 52{
		fmt.Println("num gt 15")
	}else{
		fmt.Println("num lt 15")
	}
}

func process2() int{
	return 15
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
}
