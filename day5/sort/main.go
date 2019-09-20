package main

import "fmt"

//冒泡算法
func BubbleSort(numbs *[5]int)  {
	//第一次排序，我们把最大值放在最后一个位置
	fmt.Println("排序前的值：",numbs)
	for i:=0;i<4 ;i++  {
		if (*numbs)[i] > (*numbs)[i+1]{
			temp := (*numbs)[i]
			(*numbs)[i] = (*numbs)[i+1]
			(*numbs)[i+1]= temp
		}
	}
	fmt.Println("第1次排序后的值：",numbs)

	//第二次排序，由于我们已经定位好了最大值，所以我们可以减少一次排序操作，所以就会减少一次循环操作
	for i:=0;i<3 ;i++  {
		if (*numbs)[i] > (*numbs)[i+1]{
			temp := (*numbs)[i]
			(*numbs)[i] = (*numbs)[i+1]
			(*numbs)[i+1]= temp
		}
	}
	fmt.Println("第2次排序后的值：",numbs)

	//第三次排序，已经定位好了2个最大值，所以又可以减少一次循环
	for i:=0;i<2 ;i++  {
		if (*numbs)[i] > (*numbs)[i+1]{
			temp := (*numbs)[i]
			(*numbs)[i] = (*numbs)[i+1]
			(*numbs)[i+1]= temp
		}
	}
	fmt.Println("第3次排序后的值：",numbs)
	//第三次排序，已经定位好了2个最大值，所以又可以减少一次循环
	for i:=0;i<1 ;i++  {
		if (*numbs)[i] > (*numbs)[i+1]{
			temp := (*numbs)[i]
			(*numbs)[i] = (*numbs)[i+1]
			(*numbs)[i+1]= temp
		}
	}
	fmt.Println("第4次排序后的值：",numbs)


}
//上面是逐步来进行的排序操作，现在我们就可以简化一下
func BubbleSort2(numbs *[5]int)  {
	for i:=0;i<len(numbs)-1 ;i++  {
		for j:=0;j<len(numbs)-1-1 ;j++  {
			if (*numbs)[j] > (*numbs)[j+1]{
				/*
				常规的是需要定义一个临时变量来记录最大值，下面提供了一个更简单的写法。
				temp := (*numbs)[j]
				(*numbs)[j] = (*numbs)[j+1]
				(*numbs)[j+1]= temp
				*/
				(*numbs)[j],(*numbs)[j+1] = (*numbs)[j+1],(*numbs)[j]

			}
		}
	}
}


func main() {
	intArray := [5]int{ 14,23,89,10,45}
	//BubbleSort(&intArray)
	BubbleSort2(&intArray)
	fmt.Println(intArray)
}
