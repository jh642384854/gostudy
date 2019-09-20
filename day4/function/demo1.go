package main

import "fmt"

func Test1(a int) int {
	a = 10
	return a
}

func Test2(a *int) {
	*a = 15
}

func Test3()(min int,max int)  {
	min = 1
	max = 2
	return
}

func Test4()(int,int)  {
	return  3,5
}
// 多参数方法
func ManyParams(a ...int) int {
	if(len(a) == 0){
		return 0
	}
	min := a[0]
	for _,v := range  a{
		if v< min{
			min = v
		}
	}
	return  min
}

var min,max int

func main() {

	num := 1
	Test1(num)
	fmt.Println(num);

	Test2(&num)
	fmt.Println(num);


	min,max = Test3()
	fmt.Println(min,max)

	var num1,num2 int
	num1,num2 = Test4()//这样赋值并不算是使用了这个变量。只有下面的输出的时候，使用到了这个变量，才算是用到了这个变量。
	fmt.Println(num1,num2)

	num3,num4 := Test4()
	fmt.Println(num3,num4)

	min := ManyParams(1,2,3,4,-1,5)
	fmt.Println("min : ",min)

	nums := []int{2,3,4,5}
	min2 := ManyParams(nums...)
	fmt.Println("min2:",min2)

}