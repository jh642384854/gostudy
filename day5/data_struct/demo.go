package main

import "fmt"

func main() {

	str := "3*15"
	var index int = 0
	fmt.Printf("%T \n",index)
	val := str[index:index+1]
	fmt.Println([]byte(val))
	str2 := "5"
	fmt.Println([]byte(str2)[0])
	fmt.Println(val)
}