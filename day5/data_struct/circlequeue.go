package main

import (
	"fmt"
	"errors"
	"os"
)

type CircleQueue struct {
	MaxSize int
	Data    [5]int
	Head    int
	Tail    int
}

//往队列添加数据
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("队列已经满了")
	}
	this.Data[this.Tail] = val
	this.Tail = (this.Tail + 1) % this.MaxSize
	return
}

//从队列取数据
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("队列为空")
	}
	val = this.Data[this.Head]
	this.Head = (this.Head + 1) % this.MaxSize
	return
}

//获取队列数据
func (this *CircleQueue) List() {
	fmt.Println("环形队列情况如下：")
	size := this.Size()

	tempHead := this.Head
	for i:= 0;i<size ;i++  {
		fmt.Printf("data[%d]:%d \t",tempHead,this.Data[tempHead])
		tempHead = (tempHead +1)%this.MaxSize
	}
	fmt.Println()
}

//获取所有元素值
func (this *CircleQueue) GetArray() {
	for i:=0;i<len(this.Data);i++  {
		fmt.Printf("data[%d]:%d \n",i,this.Data[i])
	}
}

//判断队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.Tail == this.Head
}

//判断队列是否已经满了
func (this *CircleQueue) IsFull() bool {
	return (this.Tail+1)%this.MaxSize == this.Head
}

//获取队列元素总数
func (this *CircleQueue) Size() int {
	return (this.Tail + this.MaxSize - this.Head) % this.MaxSize
}

func main() {
	circleQueue := &CircleQueue{
		MaxSize:5,
		Head:0,
		Tail:0,
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
			err := circleQueue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("数据添加成功")
			}
		case "get":
			val, err := circleQueue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("取出数据为：%d \n", val)
			}
		case "list":
			circleQueue.List()
		case "all":
			circleQueue.GetArray()
		case "exit":
			os.Exit(0)
		}
	}
}
