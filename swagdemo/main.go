package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	//下面这个很重要
	_ "dev/swagdemo/docs"
)
// @title Go-site Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1
// @BasePath ""

func main() {
	r := gin.Default()
	api := r.Group("/api")
	initRouter(api)
	r.Run(":8080")
}

func initRouter(r *gin.RouterGroup)  {
	r.GET("/docs/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/users",getUsers)
	r.GET("/user/:uid",getOneUser)
}

// @Summary 获取所有用户
// @Produce  json
// @Success 200 {string} json "{"RetCode":0,"UserInfo":{},"Action":"GetAllUserResponse"}"
// @Router /api/users [get]
func getUsers(c *gin.Context)  {
	users := make(map[string]interface{})
	users["zhangsan"] = "zhagnsan"
	c.JSON(http.StatusOK,users)
}

// @Summary 获取单个用户
// @Produce  json
// @Accept  json
// @Param name path string true "Name"
// @Success 200 {string} json "{"RetCode":0,"UserInfo":{},"Action":"GetOneUserResponse"}"
// @Router /api/users/{name} [get]
func getOneUser(c *gin.Context)  {
	user := make(map[string]interface{})
	user["name"] = "zhangsan"
	c.JSON(http.StatusOK,user)
}