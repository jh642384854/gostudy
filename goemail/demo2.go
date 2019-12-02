package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

/**
	具体实现可以参考：
	https://godoc.org/gopkg.in/gomail.v2

	SMTP服务器就是邮件代收发服务器，由邮件服务商提供，常见的SMTP服务器端口号：
	QQ邮箱：SMTP服务器地址：smtp.qq.com（端口：587）
	雅虎邮箱: SMTP服务器地址：smtp.yahoo.com（端口：587）
	163邮箱：SMTP服务器地址：smtp.163.com（端口：25）
	126邮箱: SMTP服务器地址：smtp.126.com（端口：25）
	新浪邮箱: SMTP服务器地址：smtp.sina.com（端口：25）


	参考链接：https://blog.csdn.net/b476178234/article/details/90171330
 */
func main() {


}

/**
	测试发现：163邮箱给QQ邮箱发送邮件会出现投递失败的问题。(会被QQ邮箱拒绝)
 */
func email163()  {
	mail := gomail.NewMessage()
	mail.SetHeader("From","jianghua0603@163.com")
	mail.SetHeader("To","15107169381@163.com")//15107169381@163.com
	mail.SetHeader("Subject","hello")
	mail.SetBody("text/html","使用Go测试发送邮件")

	d := gomail.NewDialer("smtp.163.com",25,"jianghua0603@163.com","jiangHUA19880603")
	if err := d.DialAndSend(mail); err != nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println("send email success")
	}
}
/**
	测试发现：qq给其他QQ邮箱发送邮件正常。QQ给163发送邮件也正常。
 */
func emailQQ()  {
	mail := gomail.NewMessage()
	mail.SetHeader("From","1357456562@qq.com")
	mail.SetHeader("To","jianghua0603@163.com")//15107169381@163.com
	mail.SetHeader("Subject","hello")
	mail.SetBody("text/html","使用Go测试发送邮件22")
	//smtp.qq.com 的端口可以是587或是465
	d := gomail.NewDialer("smtp.qq.com",465,"1357456562@qq.com","QQ邮箱配置的授权码")
	if err := d.DialAndSend(mail); err != nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println("send email success")
	}
}
