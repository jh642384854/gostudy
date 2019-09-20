package main

import "fmt"

func Test(){
	var intArr [5]int = [...]int{1,2,3,4,5}
	fmt.Println(intArr)

	sliceInt := intArr[1:3]
	fmt.Println(sliceInt)

	sliceInt[1] = 10
	//sliceInt[2] = 20 //虽然slice是可以动态变化的，但是这样来赋值操作的话，会出错：panic: runtime error: index out of range。

	fmt.Println()
	fmt.Println()

	fmt.Println(intArr)
	fmt.Println(sliceInt)
}

//切片的定义
func Test1(){
	var intArray = [...]int{1,3,4,5,6 }
	strings := "hello world"
	//方式一：引用数组元素
	slice1 := intArray[1:4]
	slice2 := strings[3:]  //字符串的底层就是byte数组，所以也可以进行这样的引用
	fmt.Println(slice1,slice2)

	//方式二：通过make()函数来创建
	var slice3 []int = make([]int,10)//这样这个切片的所有元素值默认都是0，不过这个默认值是依据具体的数据类型来定义的
	fmt.Println(slice3)

	//方式三：直接赋值创建
	var slice4 []int = []int{ 1,2,3,4,5}
	fmt.Println(slice4)
}

//切片的遍历
func Test2()  {
	slice1 := []int{ 1,2,3,4,5}

	//方式一：for循环
	for i:=0;i<len(slice1);i++{
		fmt.Print(slice1[i],"__")
	}

	fmt.Println()
	//方式2：for ...range循环
	for _,v := range slice1  {
		fmt.Print(v,"__")
	}
}

//切片的append()方法
func Test3()  {
	var intArray = [...]int{1,3,4,5,6 }
	slice1 := intArray[1:4]
	fmt.Println(slice1)
	//fmt.Println(slice1[3])  //在没有添加元素前，这样访问会提示越界
	slice1 = append(slice1,[]int{7,8,9}...) //append()会在底层新创建一个数组，如果这里赋值并不是slice1，则不会对slice1进行修改
	fmt.Println(slice1[3])
	fmt.Println(slice1)
}


//切片的copy()方法
func Test4()  {
	var intArray = [...]int{1,3,4,5,6 }
	slice1 := make([]int,1)
	//我们创建了一个长度为1的一个切面slice1，然后执行copy()函数，将intArray[1:4]元素复制到slice1中，可以看到intArray[1:4]这个元素不止一个，在执行copy()函数时候，也是不会报错的。
	//也就是说，copy()的第二个参数的元素个数不管是小于还是大于第一个参数，都是可以的。
	copy(slice1,intArray[1:4])
	fmt.Println(slice1)
}

func main()  {
	Test4()
}
