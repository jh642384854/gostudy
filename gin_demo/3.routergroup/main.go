package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
	路由组
 */
func main() {

	r := gin.Default()
	//路由分组1
	v1 := r.Group("/v1")
	{
		v1.GET("/user", func(c *gin.Context) {
			c.String(http.StatusOK,"v1 user")
		})

	}
	//路由分组2
	v2 := r.Group("/v2")
	{
		v2.GET("/user", func(c *gin.Context) {
			c.String(http.StatusOK,"v2 user")
		})
	}

	r.Run(":8080")
}
