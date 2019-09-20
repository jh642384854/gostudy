package controllers

import (
	"github.com/astaxie/beego"
)

// RouterController operations for Router
type RouterController struct {
	beego.Controller
}

func (r *RouterController) Get(){
	r.Ctx.WriteString(beego.AppConfig.String("appname"))
	r.Ctx.WriteString(beego.AppConfig.String("prod::httpport"))
}

func (r *RouterController) Demo(){
	r.Ctx.WriteString("diy function")
}

func (r *RouterController) DiyPost(){
	r.Ctx.WriteString("diy post")
}

func (r *RouterController) Post(){
	r.Ctx.WriteString("post")
}

