package main

import "fmt"

func test(){

	var isActive bool = false
	var isLink bool = true
	var enabled,disabled = true,false

	fmt.Println(isActive,isLink,enabled,disabled)
	fmt.Printf("isActive=%t",isActive);
}

func main() {
	test()
}
