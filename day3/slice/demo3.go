package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

func main() {

	b := []int{1,2,3}
	fmt.Println(b[:0]) //这里得到的是[]，并不是[1]这里需要特别注意。就是一个空切片

	var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}

	SortFloat64FastV1(a)
	fmt.Println(a)

	SortFloat64FastV2(a)
	fmt.Println(a)
}

/**
	第一种强制转换是先将切片数据的开始地址转换为一个较大的数组的指针，
	然后对数组指针对应的数组重新做切片操作。
	中间需要unsafe.Pointer来连接两个不同类型的指针传递。
	需要注意的是，Go语言实现中非0大小数组的长度不得超过2GB，
	因此需要针对数组元素的类型大小计算数组的最大长度范围（[]uint8最大2GB，[]uint16最大1GB，以此类推，但是[]struct{}数组的长度可以超过2GB）。
 */
func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以int方式给float64排序
	sort.Ints(b)
}

/**
	第二种转换操作是分别取到两个不同类型的切片头信息指针，任何类型的切片头部信息底层都是对应reflect.SliceHeader结构，
	然后通过更新结构体方式来更新切片信息，从而实现a对应的[]float64切片到c对应的[]int类型切片的转换。
 */
func SortFloat64FastV2(a []float64) {
	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	// 以int方式给float64排序
	sort.Ints(c)
}