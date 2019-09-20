package main

import "fmt"

/**
	选择排序
	给定 N 个项目和 L = 0 的数组，选择排序将：

	在 [L ... N-1] 范围内找出最小(或最大)项目 X 的位置，
	用第 L 项交换X，
	将下限 L 增加1并重复步骤1直到 L = N-2。
	别犹豫，让我们在上面的同一个小例子数组上尝试Selection Sort。
	在不失普遍性的情况下，我们也可以实现反向的选择排序：找到最大项目 Y 的位置并将其与最后一个项目交换。
 */

 //该方法是用来展示选择排序的的实际实现过程
func Algorithm_deduction()  {
	nums := [...]int{9, 8, 6, 4, 2, 7, 11, 3, 0, 5}

	lens := len(nums)
	maxIndex := 0;
	maxValue := nums[0]
	//第1次循环，找到最大值
	for i := 0;i<lens ;i++  {
		if nums[i] > maxValue {
			maxIndex = i
			maxValue = nums[i]
		}
	}
	if maxIndex != 0{
		nums[0],nums[maxIndex] = nums[maxIndex],nums[0]
	}
	//fmt.Println(maxIndex,maxValue)
	//fmt.Println(nums)

	//第2次循环，找到最大值，因为经过了第1轮后，已经确定了数组第一位就是最大值，所以第二次循环，就会从数组下标为1(第2个元素)开始比较
	maxValue = nums[1]
	maxIndex = 1
	for i := 1;i<lens ;i++  {
		if nums[i] > maxValue {
			maxIndex = i
			maxValue = nums[i]
		}
	}
	if maxIndex != 1{
		nums[1],nums[maxIndex] = nums[maxIndex],nums[1]
	}
	//fmt.Println(maxIndex,maxValue)

	//第3次循环，找到最大值，因为经过了第2轮后，已经确定了数组第一位就是最大值，所以第二次循环，就会从数组下标为2(第3个元素)开始比较
	maxValue = nums[2]
	maxIndex = 2
	for i := 2;i<lens ;i++  {
		if nums[i] > maxValue {
			maxIndex = i
			maxValue = nums[i]
		}
	}
	if maxIndex != 2{
		nums[2],nums[maxIndex] = nums[maxIndex],nums[2]
	}
	fmt.Println(nums)

	//第4次循环，找到最大值，因为经过了第3轮后，已经确定了数组第一位就是最大值，所以第4次循环，就会从数组下标为3(第4个元素)开始比较
	maxValue = nums[3]
	maxIndex = 3
	for i := 3;i<lens ;i++  {
		if nums[i] > maxValue {
			maxIndex = i
			maxValue = nums[i]
		}
	}
	if maxIndex != 3{
		nums[3],nums[maxIndex] = nums[maxIndex],nums[3]
	}

	fmt.Println(maxIndex,maxValue)
	fmt.Println(nums)

	//第5次循环，找到最大值，因为经过了第4轮后，已经确定了数组第一位就是最大值，所以第5次循环，就会从数组下标为4(第5个元素)开始比较
	maxValue = nums[4]
	maxIndex = 4
	for i := 4;i<lens ;i++  {
		if nums[i] > maxValue {
			maxIndex = i
			maxValue = nums[i]
		}
	}
	if maxIndex != 4{
		nums[4],nums[maxIndex] = nums[maxIndex],nums[4]
	}

	fmt.Println(maxIndex,maxValue)
	fmt.Println(nums)

	//第6次循环，找到最大值，因为经过了第5轮后，已经确定了数组第一位就是最大值，所以第6次循环，就会从数组下标为5(第6个元素)开始比较
	maxValue = nums[5]
	maxIndex = 5
	for i := 5;i<lens ;i++  {
		if nums[i] > maxValue {
			maxIndex = i
			maxValue = nums[i]
		}
	}
	if maxIndex != 5{
		nums[5],nums[maxIndex] = nums[maxIndex],nums[5]
	}

	fmt.Println(maxIndex,maxValue)
	fmt.Println(nums)

	//第7次循环，找到最大值，因为经过了第6轮后，已经确定了数组第一位就是最大值，所以第7次循环，就会从数组下标为6(第7个元素)开始比较
	maxValue = nums[6]
	maxIndex = 6
	for i := 6;i<lens ;i++  {
		if nums[i] > maxValue {
			maxIndex = i
			maxValue = nums[i]
		}
	}
	if maxIndex != 6{
		nums[6],nums[maxIndex] = nums[maxIndex],nums[6]
	}

	fmt.Println(maxIndex,maxValue)
	fmt.Println(nums)

	//第8次循环，找到最大值，因为经过了第7轮后，已经确定了数组第一位就是最大值，所以第8次循环，就会从数组下标为7(第8个元素)开始比较
	maxValue = nums[7]
	maxIndex = 7
	for i := 7;i<lens ;i++  {
		if nums[i] > maxValue {
			maxIndex = i
			maxValue = nums[i]
		}
	}
	fmt.Println(maxIndex,maxValue)
	if maxIndex != 7{
		nums[7],nums[maxIndex] = nums[maxIndex],nums[7]
	}
	fmt.Println(nums)

	//第9次循环，找到最大值，因为经过了第7轮后，已经确定了数组第一位就是最大值，所以第9次循环，就会从数组下标为8(第9个元素)开始比较
	maxValue = nums[8]
	maxIndex = 8
	for i := 8;i<lens ;i++  {
		if nums[i] > maxValue {
			maxIndex = i
			maxValue = nums[i]
		}
	}
	fmt.Println(maxIndex,maxValue)
	if maxIndex != 8{
		nums[8],nums[maxIndex] = nums[maxIndex],nums[8]
	}
	fmt.Println(nums)
}
 
 
func main() {
	//Algorithm_deduction()
	nums := [...]int{9, 8, 6, 4, 2, 7, 11, 3, 0, 5}
	lens := len(nums)
	maxValue := 0
	maxIndex := 0
	for j := 0;j<lens-1 ;j++  {
		maxValue = nums[j]  //记录当前数组的最大值
		maxIndex = j        //记录当前数组最大值的下标
		//比较元素从第二个元素开始
		for i := j+1;i<lens ;i++  {
			//这里是按照从大到小的排序，如果要按照从小到大，只需要把下面的比较符号修改一下即可，改成小于号就行
			if nums[i] > maxValue {
				maxIndex = i
				maxValue = nums[i]
			}
		}
		if maxIndex != j{
			nums[j],nums[maxIndex] = nums[maxIndex],nums[j]
		}
	}
	fmt.Println(nums)
}
