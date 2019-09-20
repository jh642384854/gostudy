package main

import "fmt"

func test(){

	//最常见的格式
	var var1 int = 10
	fmt.Println(var1)

	//不使用var关键字声明
	var14 := 14
	fmt.Printf("var14=%d",var14)
	fmt.Println()

	// 变量定义2
	var var2,var3,var4 int
	var2 = 15
	var3 = 13
	var4 = 25
	fmt.Printf("var2=%d,var3=%d,var4=%d",var2,var3,var4)
	fmt.Println()

	//变量定义形式3
	var var5,var6,var7 int = 1,2,3
	fmt.Printf("var5=%d,var6=%d,var7=%d",var5,var6,var7)
	fmt.Println()

	//上面的形式可以简化为
	var8,var9,var10 := 4,5,6
	fmt.Printf("var8=%d,var9=%d,var10=%d",var8,var9,var10)
	fmt.Println()

	var (
		var111 = 111
		var222 = 222
	)
	fmt.Printf("var111=%d,var222=%d",var111,var222)
	fmt.Println()


	// _特殊符号的使用(任何赋予它的值都会被丢弃)
	_,var11 := 1,2
	fmt.Println(var11)

	//声明变量未使用
	var var13 int
	fmt.Printf("var13=%d",var13)
}

func main() {
	test()
}
