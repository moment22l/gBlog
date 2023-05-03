package jwts

import (
	"errors"
	"fmt"
	"gBlog/global"

	"github.com/dgrijalva/jwt-go/v4"
)

// ParseToken 解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	// 密钥
	MySecret = []byte(global.Conf.Jwt.Secret)
	// 调用解析方法解析token到结构体中
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	// 如果token.Claims的类型能断言成CustomClaims且token有效, 则说明解析成功
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
