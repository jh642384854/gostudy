package main

import "fmt"

//面向过程的方式来实现的收支记录的小项目

func main() {
	key := ""
	loop := true
	flag := false

	money := 0
	balnce := 10000
	note := ""
	detail := "收支\t账户金额\t收支金额\t说    明"

	exit := ""

	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4)：")

		fmt.Scanln(&key)

		switch key {
		case "1":
			if flag {
				fmt.Println(detail)
			}else{
				fmt.Println("当前没有收支明细... 来一笔吧!")
			}
		case "2":
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			if money <= 0{
				fmt.Println("收入金额输入有误，请重新输入")
				break
			}else{
				balnce += money
			}
			fmt.Println("本次收入金额说明:")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n收入\t%v\t%v\t%v",money,balnce,note)
			flag = true
		case "3":
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			if money <= 0{
				fmt.Println("支出金额输入有误，请重新输入")
				break
			}else if money > balnce{
				fmt.Println("支出金额输入有误，请重新输入")
				break
			}else {
				balnce -= money
			}
			fmt.Println("本次支出金额说明:")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n支出\t%v\t%v\t%v",money,balnce,note)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗? y/n")
			for  {
				fmt.Scanln(&exit)
				if exit == "y" || exit == "n"{
					break
				}else{
					fmt.Println("输入有误，只能输入 y/n")
					break
				}
			}
			if exit == "y"{
				fmt.Println("退出软件")
				loop = false
			}
		default:
			fmt.Println("非法输入")
		}
		if !loop {
			break
		}
	}
	fmt.Println("退出系统")
}
