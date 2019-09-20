package main

import (
	"fmt"
)

type Integer int

type Student struct {
	Username string
	Age int
}

type Stu Student

func (integer *Integer) change(){
	*integer = *integer + 1
}

func (stu *Stu) String() string{
	str := fmt.Sprintf("Stu : Username = [%v],Age = [%v]",stu.Username,stu.Age)
	return str
}

func (stu Student) String() string{
	str := fmt.Sprintf("Student : Username = [%v],Age = [%v]",stu.Username,stu.Age)
	return str
}
/*
func (stu *Student) String() string{
	str := fmt.Sprintf("Student point : Username = [%v],Age = [%v]",stu.Username,stu.Age)
	return str
}*/


func NewStudent(name string,age int) *Student {
	return &Student{
		Username:name,
		Age:age,
	}
}
//下面这个方法并没有用到指针，和上面的方法的区别
func NewStudent2(name string,age int) Student {
	return Student{
		Username:name,
		Age:age,
	}
}

//结构体的引用方式。
func Test1()  {
	//方式1
	var stu1 Student
	stu1.Username = "zhangsan"
	stu1.Age = 10

	//方式2  方式3、方式4和这个方式所实现的效果是一样的
	var stu2 *Student = new(Student)
	stu2.Username = "lisi"
	stu2.Age = 25

	//方式3
	var stu3 *Student = &Student{
		Username:"wangwu",
		Age:15,
	}

	//方式4
	stu4 := new(Student)
	stu4.Username = "zhaoliu"
	stu4.Age = 18

	//方式5
	stu5 := Student{
		Username:"stu5",
		Age:15,
	}

	fmt.Println(stu1,stu2,stu3,stu4,stu5)
}

func Test2()  {
	stu := NewStudent("wangmazi",19)
	stu.Username = "sun"
	fmt.Println(stu)

	stu2 := NewStudent2("wang",20)
	stu2.Username = "kfjdk"
	fmt.Println(stu2)
}

func Test3()  {
	stu := Stu{"jfkdj",25}
	fmt.Println(stu)
}

//结构体的继承
type A struct {
	Name string
	Age int
	intro string //故意写的小写字母
}

type B struct {
	A  //这个结构体继承A结构体
	Name string //定义和A结构体同名的字段Name
}

func (a *A) Func1(){
	fmt.Println("A Func1()")
}

func (a *A) func2(){
	fmt.Println("A func2()")
}

func (b *B) Func1(){
	fmt.Println("B Func1()")
}

func Test4()  {
	//创建A的结构体对象
	a := &A{
		Name:"a name",
		Age :14,
		intro:"a intro",
	}
	fmt.Println(*a)
	a.Func1()
	a.Func1()

	//创建B被结构体对象
	b := &B{
		Name:"b name",
	}
	b.Age = 15
	fmt.Println(*b)
	b.Func1()
	b.A.Func1()
	b.func2()

}


func main() {
	//Test1()
	//Test2()
	//Test3()
	Test4()
}
