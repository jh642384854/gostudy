package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
	路由参数
 */
func main() {
	r := gin.Default()
	//请求示例：http://localhost:8080/zhangsan/987fbc97-4bed-5078-9f07-9141ba07c9f3
	r.GET("/:username/:uid", func(c *gin.Context) {
		var person Person
		//ShouldBindUri()需要传递地址
		if err := c.ShouldBindUri(&person); err != nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"msg":err,
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"name":person.Name,
			"id":person.ID,
		})
	})

	r.Run(":8080")
}

//定义路由参数的名称及其约定规则
type Person struct {
	ID string `uri:"uid" binding:"required,uuid"`
	Name string `uri:"username" binding:"required"`
}


