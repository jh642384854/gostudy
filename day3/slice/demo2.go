package main

import "fmt"

func AddSlice()  {
	var a = []int{1,2,3}
	a = append(a,4)       //在结尾添加一个元素
	a = append(a,[]int{6,7,8}...)  //在结尾添加一个切片
	/**
		在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。
		因此，从切片的开头添加元素的性能一般要比从尾部追加元素的性能差很多。
	 */
	a = append([]int{0},a...)      //在开头添加一个元素
	a = append([]int{-2,-1},a...)  //在开头添加一个切片

	fmt.Println(a)

	//向切片中间添加元素
	//方式一：建立临时切片方式
	var b = []int{9,8,7,6,5,4}
	b = append(b[:3],append([]int{10},a[3:]...)...) //就是在切片的索引为3的下标后面添加一个元素
	//能够添加一个元素，就可以添加多个元素
	b = append(b[:5],append([]int{-1,-2,-3},b[5:]...)...)
	fmt.Println(b)

	//方式二：利用copy()和append()函数相结合方式
	//以上向切片中间添加元素都用到了一个临时的变量，[]int{10}和[]int{-1,-2,-3}
	//可以用copy和append组合可以避免创建中间的临时切片，同样是完成添加元素的操作
	//①、向切片中间添加一个元素
	fmt.Println()
	var c = []int{9,8,7,6,5,4}
	c = append(c,0) //切片扩展1个空间
	fmt.Println(c)
	i := 3
	fmt.Println("c[i+1:]",c[i+1:],"c[i:]:",c[i:])
	copy(c[i+1:],c[i:])      //copy(dst, src []Type)函数会将src的元素逐个覆盖dst里面的元素，src的下标元素会替换相应dst元素下标的值。会影响原切片内容。c[i:]向后移动1个位置
	c[i] = 100
	fmt.Println(c)           //设置新添加的元素
	fmt.Println()


	//②、向切片中间添加多个元素
	var d = []int{9,8,7,6,5,4}
	var add = []int{-1,-2,-3}
	d = append(d,add...)        // 为x切片扩展足够的空间
	copy(d[i+len(add):],d[i:])  // a[i:]向后移动len(x)个位置
	copy(d[i:],add)             // 复制新添加的切片
	fmt.Println(d)
}

func DeleteSlice()  {
	var a = []int{9,8,7,6,5,4}

	//①、删除尾部元素
	a = a[:len(a)-1] //删除尾部一个元素
	fmt.Println(a)
	N := 2
	a = a[:len(a)-N] //删除尾部的N个元素
	fmt.Println(a)

	fmt.Println()

	//②、删除头部元素
	var b = []int{9,8,7,6,5,4}
	b = b[1:]          //删除开头1个元素
	fmt.Println(b)
	b = b[N:]          //删除开头N个元素
	fmt.Println(b)

	fmt.Println()

	//③删除开头的元素也可以不移动数据指针，但是将后面的数据向开头移动。可以用append原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）
	var c = []int{9,8,7,6,5,4}
	c = append(c[:0],c[1:]...)  // 删除开头1个元素
	fmt.Println(c)
	c = append(c[:0],c[N:]...)  // 删除开头N个元素
	fmt.Println(c)

	//使用copy()函数来删除开头元素
	var d = []int{9,8,7,6,5,4}
	d = d[:copy(d,d[1:])]   // 删除开头1个元素
	fmt.Println(d)
	d = d[:copy(d,d[N:])]  // 删除开头N个元素
	fmt.Println(d)

	fmt.Println()

	//④、删除中间元素e
	var e = []int{9,8,7,6,5,4}
	index := 2
	e = append(e[:index],e[index+1:]...) //删除中间一个元素
	fmt.Println(e)
	e = append(e[:index],e[index+N:]...) //删除中间N个元素
	fmt.Println(e)

	fmt.Println()

	//或是借助copy()函数来实现
	var f = []int{9,8,7,6,5,4}
	f = f[:index+copy(f[index:],f[index+1:])] //删除中间一个元素
	fmt.Println(f)
	f = f[:index+copy(f[index:],f[index+N:])] //删除中间N个元素
	fmt.Println(f)
}

func main() {
	DeleteSlice()
}