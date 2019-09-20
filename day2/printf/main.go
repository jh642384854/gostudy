package main

import "fmt"

func test(){
	var var1 int = 15
	// %v输出  以默认方式打印变量值
	fmt.Printf("var1 v=%v",var1)
	fmt.Println()
	// %d输出
	fmt.Printf("var1 d=%d",var1)
	fmt.Println()
	// %+d 带符号的整型
	fmt.Printf("var1 +d=%+d",var1)
	fmt.Println()
	// %q 打印单引号
	fmt.Printf("var1 q=%q",var1)
	fmt.Println()
	// %o 不带零的八进制
	fmt.Printf("var1 o=%o",var1)
	fmt.Println()

	// %#o 带零的八进制
	fmt.Printf("var1 #o=%#o",var1)
	fmt.Println()

	// %x 小写的十六进制
	fmt.Printf("var1 x=%x",var1)
	fmt.Println()

	// %#x 带0x的十六进制
	fmt.Printf("var1 #x=%#x",var1)
	fmt.Println()

	// %X 大写的十六进制
	fmt.Printf("var1 X=%X",var1)
	fmt.Println()

	// %U 打印Unicode字符
	fmt.Printf("var1 U=%U",var1)
	fmt.Println()

	// %#U 打印带字符的Unicode
	fmt.Printf("var1 #U=%#U",var1)
	fmt.Println()

	// %b打印整型的二进制
	fmt.Printf("var1 b=%b",var1)
	fmt.Println()

	// %T输出类型
	fmt.Printf("var1 T=%T",var1)
	fmt.Println()

	fmt.Printf("|%5d|", 1)
	fmt.Println()

	fmt.Printf("|%05d|", 1)
	fmt.Println()

	fmt.Printf("|%5d|", 1234567)
	fmt.Println()


	fmt.Printf("|%-5d|", 1)
	fmt.Println()

	fmt.Printf("|%-5d|", 1234567)
	fmt.Println()

	var var2 float32 = 3.14
	fmt.Printf("var2 f=%f",var2)
	fmt.Println()

	fmt.Printf("var2 .3f=%.3f",var2)
	fmt.Println()

	fmt.Printf("var2 e=%e",var2)
	fmt.Println()

	fmt.Printf("var2 g=%g",var2)
	fmt.Println()

	fmt.Printf("var2 .3g=%.3g",var2)
	fmt.Println()

	str := "hello string"
	fmt.Printf("str 5s=%5s","str")
	fmt.Println()

	fmt.Printf("str -5s=%-5s","str")
	fmt.Println()

	fmt.Printf("str .5s=%.5s",str)
	fmt.Println()

	fmt.Printf("str 5.7s=%5.7s",str)
	fmt.Println()

	fmt.Printf("str -5.7s=%-5.7s",str)
	fmt.Println()

	fmt.Printf("str 05s=%05s",str)
	fmt.Println()

	boo := false
	fmt.Printf("str t=%t",boo)
	fmt.Println()
}

func main() {
	test()
}
