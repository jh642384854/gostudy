package main
//空接口使用
import "fmt"

//定义一个空接口   空接口可以赋值给任意数据类型
type Ginterface interface {

}

func main() {
	var ginterface Ginterface
	ints := 10
	ginterface = ints  //将一个int类型的数据赋值给空接口Ginterface
	fmt.Println(ginterface)

	strings := 15
	ginterface = strings //又将一个string类型的数据赋值给空接口Ginterface

	fmt.Println(ginterface)

	var nums interface{}  //定义一个空接口数据类型
	nums = false          //然后把bool类型的数据赋值给这个空接口数据类型

	fmt.Println(nums)
}