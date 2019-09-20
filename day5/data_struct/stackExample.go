package main

import (
	"errors"
	"fmt"
	"strconv"
)

/**
	简单计算器的实现(只能实现加减乘除的运算，不能处理带小括号的)
	实现思路：
	1.创建两个栈，numStack(用来存放数字)，operatorStack(用来存放运算操作符)
	2.定义一个计算表达式(exp),是一个字符串格式
	3.逐个扫描定义的计算表达式字符串
		①、如果是数字，就直接丢进numStack
		②、如果是运算操作符
			Ⅰ：如果operatorStack是一个空栈，就直接入栈
			Ⅱ：如果不是一个空栈，又要区分如下情况：
				情形一：如果发现operatorStack栈顶的运算符的优先级大于等于当前准备入栈的运算符的优先级，就从
						operatorStack栈中pop出来，并从numStack栈也pop出两个数，并进行运算，将运算后的结果重新丢入numStack栈中
						符号在丢入operatorStack栈中。
				情形二：如果优先级不大于，则该运算符就直接入栈

	4.扫描完计算表达式以后，就需要依次从operatorStack栈中取出符号，在从numStack栈中取出两个数据，然后进行运算
	  将运算的结果在入栈，直到符号栈为空
	5.取出最后numStack栈的数据，就是计算表达式的最后值。

 */

func main() {
	//定义计算表达式
	exp := "800+3*6-1" //36
	//定义operatorStack栈
	operatorStack := &CalStack{
		MaxSize:20,
		Top:-1,
	}
	//numStack栈
	numStack := &CalStack{
		MaxSize:20,
		Top:-1,
	}
	//逐个扫描计算表达式。
	index := 0 //定义扫描计算表达式的索引
	num1 := 0  //定义因为运算符优先级较高要进行参与运算的数。从numStack栈弹出的数
	num2 := 0  //定义因为运算符优先级较高要进行参与运算的数。从numStack栈弹出的数
	res := 0   //定义因为运算符优先级较高要进行参与运算，把运算结果进行保存的变量
	operator := 0
	strNum := ""
	//下面的这步，主要用来扫描计算表达式，把里面的字符分别压入numStack栈和operatorStack栈中
	for {
		//获取每个元素值
		val := exp[index:index+1] //这里获取的是单个字符，这里需要将字符转换为ASCII来进行判断。因为对于+、-、*、/这几个符号，都可以通过ASCII来做区分
		valAscii := int([]byte(val)[0]) //注意这里的转换，[]byte(val)是将val转换为二进制切片，[]byte(val)[0]取切片的第一个元素，int([]byte(val)[0])将数据转换为int类型

		if isOperator(valAscii){
			//如果是运算符，根据上面的思路来写判断条件
			//首先判断operatorStack栈是否为空亚茹
			if operatorStack.Top == -1{
				operatorStack.PushVal(valAscii)
			}else{
				//如果operatorStack栈不为空，则需要判断运算符的优先级
				//取得operatorStack栈内的最顶部的值，比较运算符的优先级
				topOperatorPriority := priority(operatorStack.Data[operatorStack.Top])
				if  topOperatorPriority > priority(valAscii){
					//这里如果这里的优先级较高，则需要先计算结果。怎么计算呢？
					//这就需要把这个优先级高的符号先从operatorStack栈中弹出来，并且从numStack弹出两个数参与运算，然后再把运算结果压入numStack栈中，最后再把当前的运算符也压入operatorStack栈中
					num1,_ = numStack.PopVal()
					num2,_ = numStack.PopVal()
					operator,_ = operatorStack.PopVal()
					res = calc(num1,num2,operator)
					numStack.PushVal(res)
					operatorStack.PushVal(valAscii)
				}else{
					//如果要被添加的运算符优先级比operatorStack栈顶的要低，就直接入栈
					operatorStack.PushVal(valAscii)
				}
			}
		}else{
			strNum += val
			//如果不是运算符，则入numStack栈
			//将截取的字符串转换为数字
			/*number,_ := strconv.Atoi(strNum)
			numStack.PushVal(number)*/
			if index == len(exp) - 1 {
				val, _ := strconv.ParseInt(strNum, 10, 64)
				numStack.PushVal(int(val))
			} else {
				//向 index 后面测试看看是不是运算符 [index]
				if isOperator(int([]byte(exp[index+1:index+2])[0])) {
					val, _ := strconv.ParseInt(strNum, 10, 64) //将截取的字符串转换为数字
					numStack.PushVal(int(val))
					strNum = ""
				}
			}
		}
		//直到遍历结束，就退出for循环
		if index == len(exp)-1{
			break
		}
		index ++
	}
/*
	如下只是进行测试栈内的元素是否正确*/
	numStack.ListVal()
	fmt.Println()
	operatorStack.ListVal()

	//然后提取结果
	for  {
		if operatorStack.Top == -1{
			break
		}
		num1,_ = numStack.PopVal()
		num2,_ = numStack.PopVal()
		operator,_ = operatorStack.PopVal()
		res = calc(num1,num2,operator)
		numStack.PushVal(res)
	}
	result,_ := numStack.PopVal()
	fmt.Printf("最终计算的结果为：%d \n",result)
}


//判断是否是操作符
/**
	43  => +
	45  => -
	42  => *
	47  => /
 */
func isOperator(val int) bool  {
	if val == 43 || val == 45 || val == 42 || val == 47{
		return true
	}else {
		return false
	}
}

//判断运算符的优先级   *和/优先级高，设置为1，+和-优先级较低，设置为0。这里的规则完全根据自己的约定来设置
func priority(operator int) int {
	if operator == 42 || operator == 47{
		return 1
	}else{
		return 0
	}
}

//执行运算
func calc(num1 int,num2 int,operator int) int  {
	res := 0
	switch operator {
	case 43:
		res = num2 +num1
	case 45:
		res = num2 - num1
	case 42:
		res = num2*num1
	case 47:
		res = num2/num1
	default:
		fmt.Println("运算符不合法")
	}
	return res
}

//定义一个栈的结构体
type CalStack struct {
	MaxSize int
	Top int
	Data [5]int
}

// 入栈操作
func (this *CalStack)  PushVal(val int) (err error)  {
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
func (this *CalStack)  PopVal() (val int,err error) {
	//判断当前栈是否已经为空了
	if this.Top == -1{
		return -1,errors.New("stack empty")
	}
	val = this.Data[this.Top]
	this.Top --
	return val,nil
}

//遍历栈内容
func (this *CalStack)  ListVal()  {
	//判断当前栈是否已经为空了
	if this.Top == -1{
		fmt.Println("stack empty")
		return
	}
	for i:= this.Top;i >=0 ;i--  {
		fmt.Printf("data[%d]=%d \n",i,this.Data[i])
	}
}

