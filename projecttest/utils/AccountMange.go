package utils

import "fmt"

type AccountMange struct {
	key string
	loop bool
	flag bool
	money int
	balnce int
	note string
	detail string
	exit string
}
//工厂方法用来创建实例
func NewAccountMange() *AccountMange{
	return &AccountMange{
		key:"",
		loop:true,
		flag:false,
		money:0,
		balnce:10000,
		note:"",
		detail:"收支\t账户金额\t收支金额\t说    明",
		exit:"",
	}
}

func (this *AccountMange)  MainMenu() {
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4)：")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.show()
		case "2":
			this.incre()
		case "3":
			this.pay()
		case "4":
			this.quit()
		default:
			fmt.Println("非法输入")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("退出系统")
}

func (this *AccountMange)  show() {
	if this.flag {
		fmt.Println(this.detail)
	}else{
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
}

func (this *AccountMange)  incre() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	if this.money <= 0{
		fmt.Println("收入金额输入有误，请重新输入")
	}else{
		this.balnce += this.money
	}
	fmt.Println("本次收入金额说明:")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n收入\t%v\t%v\t%v",this.money,this.balnce,this.note)
	this.flag = true
}

func (this *AccountMange)  pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	if this.money <= 0{
		fmt.Println("支出金额输入有误，请重新输入")
	}else if this.money > this.balnce{
		fmt.Println("支出金额输入有误，请重新输入")
	}else {
		this.balnce -= this.money
	}
	fmt.Println("本次支出金额说明:")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n支出\t%v\t%v\t%v",this.money,this.balnce,this.note)
	this.flag = true
}

func (this *AccountMange) quit() {
	fmt.Println("你确定要退出吗? y/n")
	for  {
		fmt.Scanln(&this.exit)
		if this.exit == "y" || this.exit == "n"{
			break
		}else{
			fmt.Println("输入有误，只能输入 y/n")
			break
		}
	}
	if this.exit == "y"{
		fmt.Println("退出软件")
		this.loop = false
	}
}