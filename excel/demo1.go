package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)
/**
详细文档请查看 https://xuri.me/excelize/zh-hans/
 */

func main() {
	xlsx,err := excelize.OpenFile("demo.xlsx")
	if err != nil{
		fmt.Println(err)
		return
	}
	cell,err := xlsx.GetCellValue("Sheet1","B2")
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(cell);
}
