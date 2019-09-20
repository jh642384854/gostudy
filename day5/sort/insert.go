package main

import "fmt"

/**
	插入排序(位置对调)基本思想：
	将整个数组a分为有序和无序的两个部分。
	前者在左边，后者在右边。开始有序的部分只有a[0] , 其余都属于无序的部分。
	每次取出无序部分的第一个（最左边）元素，把它加入有序部分。
	假设插入合适的位置p，则原p位置及其后面的有序部分元素都向右移动一个位置，有序的部分即增加了一个元素。
	一直做下去，直到无序的部分没有元素。
	原文：https://blog.csdn.net/qq_41045071/article/details/81053924

	插入排序法分析：
	外循环执行N-1次，这很明显。
	但内循环执行的次数取决于输入：
	①、在最好的情况下，数组已经排序并且（a [j]> X）总是为假所以不需要移位数据，并且内部循环运行在O（1），
	②、在最坏的情况下，数组被反向排序并且（a [j]> X）始终为真插入始终发生在数组的前端，并且内部循环以O（N）运行。
	因此，最佳情况时间是O(N × 1) = O(N) ，最坏情况时间是O(N × N) = O(N2).
 */

func Algorithm_deduction2()  {
	nums := [...]int{3, 8, 6, 4}

	// 3 和8 比较
	insertVal := nums[1]  //从第二个元素开始操作
	insertIndex := 1-1    //被插入的元素位置

	//新插入的单个元素要和已经被插入的新数组的各个元素都要进行比较，找到最合适的位置

	for insertIndex >= 0 && nums[insertIndex] < insertVal {
		nums[insertIndex +1] = nums[insertIndex]
		insertIndex -- //这里要执行这个操作，因为新插入的
	}

	if insertIndex +1 != 1{
		nums[insertIndex +1 ] = insertVal
	}

	insertVal = nums[2]  //从第二个元素开始操作
	insertIndex = 2-1    //被插入的元素位置

	//新插入的单个元素要和已经被插入的新数组的各个元素都要进行比较，找到最合适的位置

	for insertIndex >= 0 && nums[insertIndex] < insertVal {
		nums[insertIndex +1] = nums[insertIndex]
		insertIndex -- //这里要执行这个操作，因为新插入的
	}

	if insertIndex +1 != 2{
		nums[insertIndex +1 ] = insertVal
	}

	//fmt.Println(nums)

	insertVal = nums[3]  //从第二个元素开始操作
	insertIndex = 3-1    //被插入的元素位置

	//新插入的单个元素要和已经被插入的新数组的各个元素都要进行比较，找到最合适的位置

	for insertIndex >= 0 && nums[insertIndex] < insertVal {
		nums[insertIndex +1] = nums[insertIndex]
		insertIndex -- //这里要执行这个操作，因为新插入的
	}

	if insertIndex +1 != 3{
		nums[insertIndex +1 ] = insertVal
	}

	fmt.Println(nums)
}
 
func InsertOrder()  {
	nums := [...]int{3, 8, 6, 4, 2, 7, 1, 35, 0, 8}
	fmt.Println("原始值如下：",nums)


	for i:=1;i<len(nums) ;i++  {
		insertVal := nums[i]   // 待插入元素
		insertIndex := i-1    // 已经排好的序列个数

		for insertIndex >= 0 && nums[insertIndex] < insertVal  { //序列从后向前遍历，将大于待插入数的元素向后移一位
			nums[insertIndex+1] = nums[insertIndex]  //元素向后移动一位，这样数组中就会有相同的两个元素值
			insertIndex --
		}

		fmt.Println("未完成排序前的值：",nums)
		fmt.Println()
		if insertIndex +1 != i{ //这里的if语句是优化用的
			nums[insertIndex+1] = insertVal   //这里就是将值进行重复的值进行替换
		}
	}
/*
	for i:=1;i<len(nums) ;i++  {
		temp := nums[i] //待插入元素
		m := i-1 //已经排好的序列个数
		for m >=0 && temp > nums[m]{ //序列从后向前遍历，将大于待插入数的元素向后移一位
			nums[m+1] = nums[m] //元素向后移动一位
			m --
		}
		nums[m+1] = temp
	}
*/
	fmt.Println(nums)
}

func main() {
	//nums := [...]int{3, 8, 6, 4, 2, 7, 1, 35, 0, 8}

	//fmt.Println(nums)

	InsertOrder()
}