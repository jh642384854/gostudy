package main

import "fmt"

type person struct {

}

func sayHi(p *person)  {
	fmt.Println("hi")
}

func (p *person) sayHi() {
	fmt.Println("person say hi")
}

func Test1()  {
	var p *person

	if p == nil {
		fmt.Println("p eq nil")
	}

	p.sayHi()
}

func main() {
	Test1()
}
