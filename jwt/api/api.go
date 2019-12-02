package api

import (
	"dev/jwt/model"
	"dev/jwt/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 注册信息
type RegistInfo struct {
	//手机号
	Phone string `json:"phone"`
	//密码
	Pwd string `json:"pwd"`
}

//登录结果结构体
type LoginResult struct {
	Token string `json:"token"`
}
// 注册用户
func ReisterUser(c *gin.Context)  {
	var registerInfo RegistInfo
	if c.BindJSON(&registerInfo) == nil{

	}
}

//用户登录
func Login(c *gin.Context)  {
	var loginReq model.LoginReq
	if c.BindJSON(&loginReq) == nil{
		user,bl := model.LoginCheck()
		if bl{
			genrateToken(c,user)
		}else{
			c.JSON(http.StatusOK,gin.H{
				"status":-1,
				"msg":"用户密码错误",
			})
		}
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status":-1,
			"msg":"json 解析失败",
		})
	}
}

//生成令牌
func genrateToken(c *gin.Context,user model.User)  {
	j := &util.JWT{
		[]byte("newtrekWang"),
	}
	claims := util.CustomClaims{
		user.ID,
		user.Name,
		user.Phone,
		jwt.StandardClaims{
			NotBefore:int64(time.Now().Unix()-1000),//签名生效时间
			ExpiresAt:int64(time.Now().Unix()+3600),//过期时间 一小时
			Issuer:"newtrekWang",//签名的发行者
		},
	}
	token,err := j.CreateToken(claims)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"status":-1,
			"msg":err.Error(),
		})
		return
	}
	log.Print(token)
	c.JSON(http.StatusOK,gin.H{
		"status":0,
		"msg":"登录成功",
		"token":token,
	})
}