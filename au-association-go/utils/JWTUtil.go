package utils

import (
	"au-golang/model/common/response"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义过期时间
const TokenExpireDuration = time.Hour * 24

var MySecret = []byte("这是一段生成token的密钥")

// 用来决定JWT中应该存储哪些数据
type MyClaims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成token并返回

func GenToken(userId string, username string) (string, error) {
	c := MyClaims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "userFunction",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

//解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, err
	}
	//token失效
	return nil, errors.New("invalid token")
}

// 后续会携带着token进行请求接口

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//请求头或者cookie或者query中获取token
		token, ok := c.GetQuery("token")

		if !ok {
			var err error
			token, err = c.Cookie("token")
			if err != nil {
				token = c.GetHeader("token")
				if token == "" {
					response.FailWithMessage("未携带token", c)
					c.Abort()
					return
				}
			}
		}

		res, err := ParseToken(token)

		if err != nil {
			response.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		//保存当前请求信息到上下文c中

		c.Set("user_id", res.UserId)
		c.Set("username", res.Username)

		//继续执行后续的请求
		c.Next()

	}
}
