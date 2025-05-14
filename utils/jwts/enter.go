package jwts

import (
	"bloger_server/global"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// 这个比较新
//import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"user_name"`
	Role     int8   `json:"role"`
}

type MyClaims struct {
	Claims
	jwt.StandardClaims
}

// 生成token
func GenToken(claims Claims) (string, error) {
	cla := MyClaims{
		Claims: claims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(global.Config.JWT.Expire)).Unix(),
			Issuer:    global.Config.JWT.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	fmt.Println(token)
	return token.SignedString([]byte(global.Config.JWT.Secret))
}

// 解析token
func parseToken(tokenString string) (*MyClaims, error) {
	if tokenString == "" {
		return nil, errors.New("tokenString is empty")
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JWT.Secret), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, errors.New("token已过期")
		}
		if strings.Contains(err.Error(), "token is malformed") {
			return nil, errors.New("token格式错误")
		}
		if strings.Contains(err.Error(), "token is signed with an unexpected signing method") {
			return nil, errors.New("token签名方法错误")
		}
		if strings.Contains(err.Error(), "token is invalid") {
			return nil, errors.New("token无效")
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func ParseTokenByGin(c *gin.Context) (*MyClaims, error) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		token = c.GetHeader("Authorization")
	}
	if token == "" {
		token = c.Query("token")
	}
	return parseToken(token)
}
