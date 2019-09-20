package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)
/**
	中间件的定义和使用
 */
func main() {
	r := gin.New()
	r.Use(DiyLogger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example)
	})

	r.Run(":8080")
}
/**
	gin.Default()和gin.New()的区别
 */
func Test1()  {
	//r := gin.Default()
	r := gin.New()
	fmt.Println(r)
}

/**
	这个是一个自定义的中间件
 */
func DiyLogger() gin.HandlerFunc  {
	return func(c *gin.Context) {
		t := time.Now()
		//给gin.Context.Keys设置值。这个值会在每个请求都会获取
		c.Set("example","123456")
		c.Next()
		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}