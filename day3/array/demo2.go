package main

import "fmt"

func FunArrDefinded()  {
	/**
		第一种方式是定义一个数组变量的最基本的方式，数组的长度明确指定，数组中的每个元素都以零值初始化。
	 */
	var a [3]int
	/**
		第二种方式定义数组，可以在定义的时候顺序指定全部元素的初始化值，数组的长度根据初始化元素的数目自动计算。
	 */
	var b = [...]int{1,2,3}
	/**
		第三种方式是以索引的方式来初始化数组的元素，因此元素的初始化值出现顺序比较随意。
		这种初始化方式和map[int]Type类型的初始化语法类似。
		数组的长度以出现的最大的索引为准，没有明确初始化的元素依然用0值初始化。
	 */
	var c = [...]int{2:33,1:12}
	/**
		第四种方式是混合了第二种和第三种的初始化方式，前面两个元素采用顺序初始化，
		第三第四个元素零值初始化，第五个元素通过索引初始化，
		最后一个元素跟在前面的第五个元素之后采用顺序初始化。
	 */
	var d = [...]int{1,2,7:5,4:5,9:33,6}

	fmt.Println("a [3]int:",a)
	fmt.Println("b  [...]int{1,2,3}",b)
	fmt.Println("c  [...]int{2:3,1:2}",c)
	fmt.Println("d  [...]int{1,2,4:5,6}",d)
}

func ArrPrint()  {
	var a = [...]int{1,2,3}
	var b = &a            //b 是指向数组的指针
	b[1] = 5
	fmt.Println(a[0],a[1])
	fmt.Println(b[0],b[1]) //通过数组指针访问数组元素的方式和数组类似


	fmt.Println()

	for i, v := range b {
		fmt.Println(i,v)
	}

	fmt.Println()
	for i, v := range a {
		fmt.Println(i,v)
	}
}

//空数组定义
func EmptyArray()  {
	var a [0]int
	var b = [0]int{}
	var c = [...]int{}
	fmt.Println(a,b,c)
}

//空数组的应用：管道同步操作
func EmptyArrayApply()  {
	c1 :=make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	fmt.Println(<- c1)

	//更加合适的方法:用无类型的匿名结构体代替
	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{}  //struct{}部分是类型, {}表示对应的结构体值
	}()
	fmt.Println(<- c2)
}


func main()  {
	//ArrPrint()
	//EmptyArray()
	EmptyArrayApply()
}