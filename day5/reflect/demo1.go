package main

import (
	"fmt"
	"reflect"
)

func TestFun1()  {
	//var data interface{}
	var data = 10
	rType := reflect.TypeOf(data)
	fmt.Printf("rType value:%v,rType type:%T \n",rType,rType)

	rValue := reflect.ValueOf(data)
	fmt.Printf("rType value:%v,rType type:%T \n",rValue,rValue)


	//将rValue转换为interface{}
	iValue := rValue.Interface()
	//将interface{}通过断言转换为需要的类型
	num2 := iValue.(int)
	fmt.Println("num2 :",num2)
}
/**
	反射可以将“接口类型变量”转换为“反射类型对象”

	大家可能对下面的案例感到疑惑，程序里没有接口类型变量啊，哪来的接口类型变量到反射类型对象的转换啊？
	事实上，reflect.ValueOf和reflect.TypeOf的参数类型都是interface{}，即空接口类型，而返回值的类型是reflect.Value和reflect.Type，中间的转换由reflect包来实现。
 */
func TestFun2()  {
	var a int = 30

	v := reflect.ValueOf(a) //返回reflect.Value类型对象，值是30，注意不能用这个值直接进行加减乘除等运算，因为类型不对。
	t := reflect.TypeOf(a)  //返回reflect.Type类型对象，值为int
	fmt.Println(v)
	fmt.Println(t)
	v = reflect.ValueOf(&a)  //返回Value类型对象，值为&a，变量a的地址
	t = reflect.TypeOf(&a) //返回Type类型对象，值为*int
	fmt.Println(v)
	fmt.Println(t)
}

/**
	反射可以将“反射类型对象”转换为“接口类型变量”
 */
func TestFun3()  {

	var a int = 30
	v := reflect.ValueOf(&a)
	t := reflect.ValueOf(&a)

	fmt.Println(v)
	fmt.Println(t)

	v1 := v.Interface() //返回空接口类型
	fmt.Println("v1:",v1)
	v2 := v1.(*int) //类型断言，断定v1的type=*int(即指针int类型)

	fmt.Printf("%T,%v \n",v2,v2)

	fmt.Println(*v2) //获取指针数据，就需要在变量前面添加一个*(取值符)
}

/**
	如果要修改反射类型对象，其值必须是“addressable”

	通过反射定义一可以知道，反射对象包含了接口变量中存储的值以及类型。
	如果反射对象中包含的值是原始值，那么可以通过反射对象修改原始值；
	如果反射对象中包含的值不是原始值（反射对象包含的是副本值或指向原始值的地址），那么该反射对象是不可以修改的。
	通过CanSet函数可以判定反射对象是否可以修改。
 */

func TestFun4()  {
	var x float64 = 3.14
	v := reflect.ValueOf(x)
	//v.SetFloat(7.1) // 这里会出现“panic: reflect: reflect.Value.SetFloat using unaddressable value”
	fmt.Println("settability of v",v.CanSet()) //这里返回false，表示v不可以修改，是因为v中保存的是3.14的副本

	//换一种方式
	var y float64 = 3.14
	v2 := reflect.ValueOf(&y) //上面的反射对象v2不可以修改，是因为v2当前保存的是y的地址，而不是y的原始空间。
	fmt.Println("settability of v2",v2.CanSet())

	//继续用另外一种方式
	var z float64 = 3.14
	v3 := reflect.ValueOf(&z).Elem() //这里相比上面就多调用了reflect.Value.Elem()方法，这个方法用来获取原始值对应的反射对象
	fmt.Println("settability of v3",v3.CanSet())
	v3.SetFloat(6.15)
	fmt.Println(v3)

	//继续用另外一种方式
	var m float64 = 3.14
	v4 := reflect.ValueOf(m).Elem() //调用Elem()方法，如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。所以这里会出错
	fmt.Println("settability of v3",v4.CanSet())
	/*v4.SetFloat(9.15)
	fmt.Println(v4)*/

}


func main() {
	TestFun4()
}