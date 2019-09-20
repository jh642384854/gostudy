package controllers

import "github.com/astaxie/beego"

type AnnotationController struct {
	beego.Controller
}
// @router /annotation/demo1
func (c *AnnotationController) Demo1()  {
	c.Ctx.WriteString("AnnotationController Demo1")
}


