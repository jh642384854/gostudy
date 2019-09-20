package main

import (
	"fmt"
	"errors"
)

/**
	栈(堆栈)结构(遵循先进后出原则)
	首先要了解栈的基本构成：
	栈由栈的容量(即栈内能存放多少元素)、栈顶、栈底来构成。栈顶就是当前栈内的最大下标的元素值，栈低始终都为-1，栈底是不会变化的。
	有了上面的栈的结构了解，我们就可以定义一个结构体来标识栈
 */

 type Stack struct {
 	MaxSize int
 	Top int
 	Data [5]int
 }

// 入栈操作
func (this *Stack)  Push(val int) (err error)  {
	//判断栈是否已经满了
	if this.Top == this.MaxSize{
		fmt.Println("stack full")
		return errors.New("statck full")
	}
	this.Top++
	this.Data[this.Top] = val
	//fmt.Println(this.Data)
	return
}

// 出栈操作
func (this *Stack)  Pop() (val int,err error) {
	//判断当前栈是否已经为空了
	if this.Top == -1{
		return -1,errors.New("stack empty")
	}
	val = this.Data[this.Top]
	this.Top --
	return val,nil
}

//遍历栈内容
func (this *Stack)  List()  {
	//判断当前栈是否已经为空了
	if this.Top == -1{
		fmt.Println("stack empty")
		return
	}
	for i:= this.Top;i >=0 ;i--  {
		fmt.Printf("data[%d]=%d \n",i,this.Data[i])
	}
}


func main() {

	stack := &Stack{
		MaxSize:5,
		Top:-1,
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	stack.List()

	stack.Pop()
	fmt.Println()

	stack.List()


}