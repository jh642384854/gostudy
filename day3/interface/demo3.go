package main
//直接创建接口变量
import "fmt"
//定义两个接口
type Binterface interface {
	Bfun1()
}

type Cinterface interface {
	Cfun1()
}
//定义一个结构体
type Bstruct struct {

}

//定义两个方法，这两个方法分别是上面两个结构体中的方法，也就是说，这个结构体同时实现了多个接口。
func (b Bstruct) Bfun1()  {
	fmt.Println("Bstruct  Bfun1()")
}

func (b Bstruct) Cfun1()  {
	fmt.Println("Bstruct  Cfun1()")
}


func main() {
	bstruct := Bstruct{}
	var binterface Binterface = bstruct
	var cinterface Cinterface = bstruct
	binterface.Bfun1()
	cinterface.Cfun1()
}
