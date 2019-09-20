package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//自定义数据类型的排序实现

//1.首先定义一个结构体
type Person struct {
	Name string
	Age int
}
//2.在定义一个切片类型(为了调用sort包的Sort()方法)
type ByAge []Person

//3.为了调用sort包的Sort()方法，还需要实现三个方法
func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Less(i,j int) bool {
	return b[i].Age < b[j].Age
}

func (b ByAge) Swap(i,j int)  {
	b[i],b[j] = b[j],b[i]
}

func main() {
	var ByAgepersons ByAge
	for i := 0;i<10 ;i++  {
		person := Person{
			Name:fmt.Sprintf("user_%d", rand.Intn(100)),
			Age:rand.Intn(100),
		}
		ByAgepersons = append(ByAgepersons,person)
	}

	persons := []Person{}
	for i := 0;i<10 ;i++  {
		person := Person{
			Name:fmt.Sprintf("user_%d", rand.Intn(100)),
			Age:rand.Intn(100),
		}
		persons = append(persons,person)
	}
	sort.Sort(ByAge(persons))
	fmt.Println(persons)
/*
	persons := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	person := Person{
		Name:fmt.Sprintf("user_%d", rand.Intn(100)),
		Age:rand.Intn(100),
	}
	persons = append(persons,person)
	sort.Sort(ByAge(persons))
	fmt.Println(persons)
*/
	fmt.Println()

	sort.Sort(ByAgepersons)
	fmt.Println(ByAgepersons)
}
