package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
	静态资源处理和html渲染
 */
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("dist/*.html")//这种是通过正则表达式匹配相应的模版文件。
	//还有一个函数LoadHTMLGlob(pattern string)，这个函数是传递具体哪些可用的模版文件，如下所示
	//r.LoadHTMLFiles("dist/index.html","dist/home.html")
	//静态资源映射配置
	r.Static("/static","./dist/static")  //这个函数其实内部也是调用的下面的StaticFS()这个函数
	//r.StaticFS("/static",http.Dir("./dist/static"))
	r.StaticFile("/favicon.ico","./dist/favicon.ico")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html", gin.H{
			"title":"Main website",
		})
	})
	r.Run(":8080")
}
