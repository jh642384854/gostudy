package main

import "fmt"
/**
	接口入门示例
 */

//定义一个接口
type User interface {
	DoJob()
	Study()
}
//定义两个结构体
/**
	这里要明确一个事情：
	1.我们这里说的一个结构体实现接口，是指的这个结构体实现了这个接口的所有方法，而不是其中的某个或是某几个。
	2.除了要实现所有接口定义的方法外，所实现的方法签名和接口定义的方法签名也必须一样。(方法签名：就是定义方法的过程，具体就是方法的名称，方法所可能会需要的形参，以及方法可能会有的返回值)
	3.结构体实现接口的过程也很简单，只需要定义结构体方法的时候，方法的名称和接口中定义的方法名称保持一致即可。
 */
type Student struct {

}

type Engineer struct {

}
//为结构体绑定方法，绑定的方法名称是接口为实现的方法名称。
func (s *Student) DoJob()  {
	fmt.Println("Student DoJob")
}

func (s *Student) Study()  {
	fmt.Println("Student Study")
}

func (e *Engineer) DoJob()  {
	fmt.Println("Engineer DoJob")
}

func (e *Engineer) Study()  {
	fmt.Println("Engineer Study")
}

//定义一个结构体，通用的来调用，根据传递的实现接口类型不一样，执行的结果也就会不一样。
type People struct {

}

func (p People) Wroking(user User)  {
	user.DoJob()
	user.Study()
}

func main() {
	people := People{}
	student := &Student{}  //这里必须要写成&Student{}，因为我们在实现User接口的方法时候，传递的是指针，如func (s *Student) DoJob(){}。如果不想写&符号，那我们在定义方法实现的时候，就不写引用传递方式，改用值传递，即func (s Student) DoJob(){}
	engineer := &Engineer{}
	people.Wroking(student)

	fmt.Println()

	people.Wroking(engineer)
}