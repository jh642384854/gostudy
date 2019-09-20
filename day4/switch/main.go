package main

import "fmt"

func Test1() {
	num := 10
	switch num {
	case 1:
		fmt.Println(1)
	case 10:
		fmt.Println(10)
	default:
		fmt.Println("default")
	}
}

func Test2() {
	i := 2
	switch i {
	case 0:
		fmt.Println(0)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	}
}

//多值条件
func Test3() {
	num := 99
	switch num {
	case 98, 99:
		fmt.Println("98 or 99")
	case 100:
		fmt.Println(100)
	}
}

//case设定为条件,switch不给任何变量
func Test4() {
	num := 10
	switch {
	case num < 10:
		fmt.Println("lt 10")
	case num > 9:
		fmt.Println("gt 9")
	}
}

//如果case后面只有一行代码，可以直接写在冒号后面
func Test5() {
	num := 10
	switch {
	case num < 10 : fmt.Println("lt 10")
	case num > 9 : fmt.Println("gt 9")
	}
}

func Test6()  {
	switch num := process(); {
	case num >= 10:
		fmt.Println("egt 10")
	case num < 10:
		 fmt.Println("lt 10")
	}
}

func process() int {
	return  10
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
}
