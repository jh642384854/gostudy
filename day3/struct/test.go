package main

import "fmt"

type A1 struct {
	Name string
	Age int
	int
}

type B1 struct {
	Name string
	Score int
}

//结构体的多继承，且字段重复
type C1 struct {
	A1
	B1
	Name string
}

//有名结构体的继承
//如果一个结构体中是有名的结构体，则访问有名结构体字段的时候，必须要带上有名结构体的名字，如下面Demo2()方法中的写法：d1.a.Name = "d1"
type D1 struct {
	a A1
	b B1
}

//结构体继承传递的是指针
type E1 struct {
	*A1
	*B1
}

//结构体继承基本数据类型
type F1 struct {
	A1
	int  //在这个结构体中，不能有多个同样的匿名的基本数据类型，但是父类是可以包含同类型的基本数据类型(当然也只能包含一个同类型的基本数据类型)
}

func Demo1()  {
	var c1 C1
	c1.Name = "c1"
	fmt.Println(c1)
}

func Demo2()  {
	var d1 D1
	//d1.Name = "d1"
	d1.a.Name = "d1"
}

func Demo3()  {
	//嵌套匿名结构体后，可以在创建结构体变量(实例)时候，直接指定各个匿名结构体字段的值。
	d1 := &D1{ A1{ Name:"D1",Age:15,},B1{Name:"D1",Score:87,}, }
	fmt.Println(d1)
}


func Demo4()  {
	e1 := E1{ &A1{ Name:"E1",Age:15,},&B1{Name:"E1",Score:87,}, }
	fmt.Println(e1)//这个获取的两个元素都是地址
	fmt.Println(*e1.A1,*e1.B1)//通过*(取值符)来获取元素的值。
}

func Demo5()  {
	f1 := F1{}
	f1.Name = "f1"
	f1.Age = 42
	f1.A1.int = 789
	f1.int = 456
	fmt.Println(f1)
}

func main() {
	//Demo3()
	//Demo4()
	Demo5()
}
