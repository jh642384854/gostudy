package main

import (
	"fmt"
	"errors"
	"os"
)

/*
	单向队列
 */

type SignalQueue struct {
	MaxSize int
	data [4]int
	head int
	tail int
}
//往队列添加元素
func (this *SignalQueue) Push(val int) (err error) {
	//先判断队列是否已经满了
	if this.tail == this.MaxSize-1{
		return errors.New("队列已经满了")
	}
	this.tail++ //tail 后移
	this.data[this.tail] = val
	return
}
//从队列取出数据
func (this *SignalQueue) Pop() (val int,err error) {
	//判断队列数据是否为空
	if this.IsEmpty(){
		fmt.Println("当前队列为空")
		return
	}
	this.head++
	val = this.data[this.head]
	return val,nil
}

//查看队列的数据列表
func (this *SignalQueue) List()  {
	if this.IsEmpty(){
		fmt.Println("当前队列为空")
		return
	}
	fmt.Println("当前队列的情况如下：")
	for i:=this.head+1;i<=this.tail ;i++  {
		fmt.Printf("array[%d]:%d \t",i,this.data[i])
	}
	fmt.Println()
}

func (this *SignalQueue) GetArray()  {
	for i:=0;i<len(this.data);i++  {
		fmt.Printf("data[%d]:%d \n",i,this.data[i])
	}
}

//判断队列是否为空
func (this *SignalQueue) IsEmpty() bool {
	return this.head == this.tail
}

func main() {
	signalQueue := &SignalQueue{
		MaxSize:4,
		head:-1,
		tail:-1,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入list 表示显示队列")
		fmt.Println("4. 输入all 表示显示数组元素")
		fmt.Println("5. 输入exit 表示显示队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入要添加的元素")
			fmt.Scanln(&val)
			err := signalQueue.Push(val)
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println("数据添加成功")
			}
		case "get":
			val,err := signalQueue.Pop()
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Printf("取出数据为：%d \n",val)
			}
		case "list":
			signalQueue.List()
		case "all":
			signalQueue.GetArray()
		case "exit":
			os.Exit(0)
		}
	}
}