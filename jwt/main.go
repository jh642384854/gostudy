package main

import (
	"dev/jwt/api"
	"dev/jwt/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Writer.WriteString("jwt demo")
	})
	r.POST("/login",api.Login)

	admin := r.Group("admin").Use(util.JWTAuth())
	{
		admin.GET("userinfo", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"status":1,
				"msg":"获取用户信息",
			})
		})
	}
	r.Run(":8080")
}
