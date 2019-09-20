package main

import "fmt"

func main() {
	a := [...]string{"a","b","c"}

	fmt.Printf("T%",a)

	fmt.Println("cap:",cap(a))

	fmt.Println("len:",len(a))

	for i := 0; i<cap(a); i++  {
		fmt.Println("Array item",i,"is",a[i])
	}
	
	for i := range  a {
		fmt.Println("Array item",i,"is",a[i])
	}
}