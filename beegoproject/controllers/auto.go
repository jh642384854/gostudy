package controllers

import "github.com/astaxie/beego"

type AutoController struct {
	beego.Controller
}

func (c *AutoController) Login(){
	c.Ctx.WriteString("AutoContoller Login()")
}

func (c *AutoController) Logout(){
	c.Ctx.WriteString("AutoContoller Logout()")
}

func (c *AutoController)  UserInfo()  {
	ext := c.Ctx.Input.Param(":ext")
	var userinfo string
	switch ext {
	case "json":
		userinfo = "json data result"
	case "xml":
		userinfo = "xml data result"
	default:
		userinfo = "normal data result"
	}
	c.Ctx.WriteString(userinfo)
}
