package main

import (
	"beegoproject/controllers"
	_ "beegoproject/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.AutoRouter(&controllers.AutoController{})
	beego.Run()
}

