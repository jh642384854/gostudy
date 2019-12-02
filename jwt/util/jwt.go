package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

/**
	可以参考https://github.com/Bingjian-Zhu/gin-vue/部分功能
	https://www.jianshu.com/p/1f9915818992
	https://mojotv.cn/go/golang-jwt-auth
 */

//定义JWT 签名结构
type JWT struct {
	SigningKey []byte
}

//定义的一些常量
var (
	TokenExpired error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed error = errors.New("That's not even a token")
	TokenInvalid error = errors.New("Counldn't handle this token:")
	SignKey string = "newtrekWang"
)

// 载荷，可以加一些自己所需要的信息
type CustomClaims struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// 新创建一个jwt实例
func NewJwt() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

//获取signkey
func GetSignKey() string {
	return SignKey
}

//设置signkey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}
//生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string,error)  {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(j.SigningKey)
}

//更新token
func (j *JWT) RefreshToken(tokenString string) (string,error)  {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0,0)
	}
	token,err := jwt.ParseWithClaims(tokenString,&CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey,nil
	})
	if err != nil{
		return "",err
	}
	if claims,ok := token.Claims.(*CustomClaims);ok && token.Valid{
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1*time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "",TokenInvalid
}

//解析Token
func (j *JWT) PareseToken(tokenString string) (*CustomClaims,error)  {
	token,err := jwt.ParseWithClaims(tokenString,&CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey,nil
	})
	if err != nil{
		if ve,ok := err.(*jwt.ValidationError);ok{
			if ve.Errors & jwt.ValidationErrorMalformed != 0{
				return nil,TokenMalformed
			}else if ve.Errors & jwt.ValidationErrorExpired != 0{
				return nil,TokenExpired
			}else if ve.Errors & jwt.ValidationErrorNotValidYet != 0{
				return nil,TokenNotValidYet
			}else{
				return nil,TokenInvalid
			}
		}
	}
	if claims,ok := token.Claims.(*CustomClaims);ok && token.Valid{
		return claims,nil
	}
	return nil,TokenInvalid
}

//定义JWTAuth中间件，检查token
func JWTAuth() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token ==  "" {
			c.JSON(http.StatusOK,gin.H{
				"status":1,
				"msg":"请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		log.Print("get Token :",token)
		j := NewJwt()
		//解析token包含的信息
		claims,err := j.PareseToken(token)
		if err != nil{
			if err == TokenExpired{
				c.JSON(http.StatusOK,gin.H{
					"status":-1,
					"msg":"授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK,gin.H{
				"status":-1,
				"msg":err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("claims",claims)
	}
}