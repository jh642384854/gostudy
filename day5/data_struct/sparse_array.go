package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/**
	稀疏数组的操作
 */

var chessboard [3][5]int
var chessboardFile string

func init() {
	chessboard[1][2] = 1
	chessboard[2][3] = 2
	chessboardFile = "E:/chessboard.data"
}

//定义一个结构体，用来记录字段信息
type ValueNode struct {
	Row int
	Col int
	Val int
}

func WriteData()  {

	fmt.Println("稀疏数组原始值如下：")
	var (
		index1 = 0
		index2 = 0
	)
	for _,v := range chessboard {
		for _,v2 := range v{
			if index1 == 0{
				index2 ++
			}
			fmt.Printf("%d\t",v2)
		}
		index1 ++
		fmt.Println()
	}

	//将上面的稀疏数组进行压缩，存放到一个切片中
	var ValueNodeSice []ValueNode
	valueNode := ValueNode{
		Row:index1,
		Col:index2,
		Val:0,
	}
	ValueNodeSice = append(ValueNodeSice,valueNode)

	for i,v := range chessboard {
		for j,v2 := range v {
			if v2 != 0{
				valueNode := ValueNode{
					Row:i,
					Col:j,
					Val:v2,
				}
				ValueNodeSice = append(ValueNodeSice,valueNode)
			}
		}
	}
	fmt.Println()
	//遍历稀疏数组
	nodeStr := ""
	for i,v :=range ValueNodeSice {
		fmt.Printf("索引%d,row:%d,col:%d,val:%d \n",i,v.Row,v.Col,v.Val)
		nodeStr += fmt.Sprintf("%d %d %d\r\n",v.Row,v.Col,v.Val)
	}

	//存盘写入数据

	file,error := os.OpenFile(chessboardFile,os.O_WRONLY|os.O_CREATE,06666)
	defer  file.Close()
	if error != nil{
		fmt.Println("文件创建失败")
		return
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(nodeStr)
	writer.Flush()

	fmt.Println("稀疏数据操作完成")
}

//从稀疏数组中还原原数组
func ReadData()  {
	fileHandler,err := os.Open(chessboardFile)
	defer fileHandler.Close()
	if err != nil{
		fmt.Println("文件读取失败")
		return
	}

	index := 0
	var chessboard2 [][]int
	reader := bufio.NewReader(fileHandler)

	for {
		index++
		chstring,err := reader.ReadString('\n')
		//chstring,err := reader.ReadBytes('\n')
		if err == io.EOF{
			break
		}

		arrString := strings.Split(string(chstring)," ")
		row,_ := strconv.Atoi(arrString[0])
		col,_ := strconv.Atoi(arrString[1])
		val,_ := strconv.Atoi(strings.Replace(arrString[2],"\r\n","",-1))
		//fmt.Println(row,col,val)
		if index == 1 {
			for i:= 0;i < row; i++  {
				var coldata []int
				for j := 0; j < col; j++ {
					coldata = append(coldata, val)
				}
				chessboard2 = append(chessboard2, coldata)
			}
		}
		if index != 1{
			chessboard2[row][col] = val
		}
	}
	//循环读取还原的数据
	for _,v := range chessboard2{
		for _,j :=range v {
			fmt.Printf("%d\t",j)
		}
		fmt.Println()
	}
}

func main() {

	WriteData()

	fmt.Println()

	ReadData()

}
