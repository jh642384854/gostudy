package main
//接口的多继承
import "fmt"
//定义两个接口
type Dinterface interface {
	Dfun1()
}

type Einterface interface {
	Efun1()
}
//定义一个接口，这个接口继承Dinterface和Einterface这两个接口(这两个或多个接口不能有同名的方法，不然会报错)，所以如果一个结构体要实现Finterface这个接口，那就需要同时实现这三个接口的所有方法。
type Finterface interface {
	Dinterface
	Einterface
	Ffun1()
}
//定义一个结构体
type Fstruct struct {

}

//定义两个方法，这两个方法分别是上面两个结构体中的方法，也就是说，这个结构体同时实现了多个接口。
func (f Fstruct) Dfun1()  {
	fmt.Println("Fstruct  Dfun1()")
}

func (f Fstruct) Efun1()  {
	fmt.Println("Fstruct  Efun1()")
}

func (f Fstruct) Ffun1()  {
	fmt.Println("Fstruct  Ffun1()")
}


func main() {
	fstruct := Fstruct{}
	var finterface Finterface = fstruct
	var dinterface Dinterface = fstruct
	var einterface Einterface = fstruct

	dinterface.Dfun1()

	einterface.Efun1()
	//var cinterface Finterface = fstruct
	finterface.Dfun1()
	finterface.Efun1()
	finterface.Ffun1()
}
