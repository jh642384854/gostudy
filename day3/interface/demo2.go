package main
//基本数据类型也可以实现接口的方法
import "fmt"
//定义一个接口
type Ainterface interface {
	ATest1()
}

//定义一个结构体
type Ainter struct {
	Name string
}

//定义一个基本数据类型
type Integer int

//定义自定义结构体的方法实现
func (a Ainter) ATest1()  {
	fmt.Println("Ainter struct ATest1()")
}

//定义基本数据类型的方法实现
func (i Integer) ATest1()  {
	fmt.Println("Integer ATest1()")
}



func main() {
	ainter := Ainter{}
	var i Integer = 10

	var a1 Ainterface  //这个是一个nil值。
	fmt.Println(a1)

	var a Ainterface = ainter  // 接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)
	a.ATest1()                 // 通过接口直接调用接口的方法。
	ainter.ATest1()            // 通过结构体来调用方法

	var b Ainterface = i
	b.ATest1()
}
