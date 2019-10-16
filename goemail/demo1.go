package main

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"strings"
)

func main() {
	receiverUsers := "642384854@qq.com"
	myEmail := &EmailParam{
		ServerHost:"smtp.qq.com",
		ServerPort:465,
		FromEail:"1357456562@qq.com",
		FromPassword:"awrujjcxznjkjbff",  //vltgnhlbdwjsgehh   nzgcplfonjuxgbec   //awrujjcxznjkjbff
		ReceiverUsers:receiverUsers,
	}
	subject := "go lang send email"
	body := "go lange send email to QQ"
	InitMail(myEmail)
	SendEmail(subject,body,myEmail)
}

type EmailParam struct {
	//邮箱服务器地址。比如腾讯企业邮箱服务器smtp.exmail.qq.com
	ServerHost string
	//邮箱服务器端口，如腾讯企业邮箱的465
	ServerPort int
	//发件人邮箱地址
	FromEail string
	// 发件人密码。这里是明文形式。TODO:如何改成密文的方式？
	FromPassword string
	// 接收者的邮箱地址，多个邮箱地址用(",")隔开，不能为空
	ReceiverUsers string
	// 抄送者的邮箱地址，多个邮箱地址用(",")隔开。
	CopyUsers string
}
var email *gomail.Message
//初始化
func InitMail(ep *EmailParam)  {
	toUsers := []string{}
	email = gomail.NewMessage()
	if len(ep.ReceiverUsers) == 0{
		return
	}
	for _, val := range strings.Split(ep.ReceiverUsers,",") {
		toUsers = append(toUsers,val)
	}
	//收件人可以是多个
	email.SetHeader("To",toUsers...)
	//抄送列表
	if len(ep.CopyUsers) != 0 {
		for _, val := range strings.Split(ep.CopyUsers,",") {
			toUsers = append(toUsers,val)
		}
		email.SetHeader("Copy to",toUsers...)
	}
	//设置发件人信息
	email.SetAddressHeader("From",ep.FromEail,"")
}
//发送邮件
func SendEmail(subject,body string,ep *EmailParam)  {
	email.SetHeader("Subject",subject)
	email.SetBody("text/html",body)
	d := gomail.NewDialer(ep.ServerHost,ep.ServerPort,ep.FromEail,ep.FromPassword)
	//发生邮件
	err := d.DialAndSend(email)
	if err != nil{
		fmt.Println(err.Error())
	}
}