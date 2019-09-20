package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

var (
	modelFile = "E:/GoProjects/src/dev/casbindemo/demo1/rbac_model2.conf"
	policyFile = "E:/GoProjects/src/dev/casbindemo/demo1/rbac_policy.csv"
)

func main() {

	casbinEnforce,err := casbin.NewEnforcer(modelFile,policyFile)
	if err != nil{
		fmt.Println("casbin.NewEnforcer Error",err.Error())
	}


	r := gin.New()
	r.Use(casbinInterceptor(casbinEnforce))

	r.GET("/user",casbinHandler())
	r.POST("/user",casbinHandler())
	r.GET("/test",casbinHandler())

	r.Run(":8080")
}

func casbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.WriteString("hello"+c.Request.RequestURI)
	}
}

func casbinInterceptor(e *casbin.Enforcer) gin.HandlerFunc  {
	return func(c *gin.Context) {
		//获取请求的URL
		obj := c.Request.URL.RequestURI()
		//获取请求的方法
		act := c.Request.Method
		//获取用户角色(注意是角色，不是用户名)
		sub := "abc123"
		//判断是否有权限
		 boolVal,err := e.Enforce(sub,obj,act)
		 if err != nil{
		 	fmt.Println("casbin Enforce error",err.Error())
		 }
		 if boolVal{
		 	fmt.Println("has promission")
		 	c.Next()
		 }else{
			fmt.Println("no promission")
			c.Abort()
		 }
	}
}