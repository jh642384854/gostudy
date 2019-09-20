package main

import "fmt"

type Jhinterface interface {
	Fun1()
}

type JhStruct struct {
	
}

type JhStruct2 struct {

}

func (j *JhStruct) Fun1()  {
	fmt.Println("JhStruct Fun1()")
}

func main() {
/*
	//直接实例化一个接口对象，然后通过这个接口对象调用自己定义的方法，是无法操作的
	var jhinterface Jhinterface
	jhinterface.Fun1()
*/
	var jhinterface Jhinterface
	jhinterface = &JhStruct{}
	jhstruct := &JhStruct{}
	jhinterface.Fun1()
	jhstruct.Fun1()
}