package jwts

import (
	"gBlog/global"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// GenToken 生成token
func GenToken(user JwtPayload) (string, error) {
	// 密钥
	MySecret = []byte(global.Conf.Jwt.Secret)
	// 通过传入的payload以及默认的jwt配置信息生成声明
	claim := CustomClaims{
		JwtPayload: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Conf.Jwt.Expires))),
			Issuer:    global.Conf.Jwt.Issuer,
		},
	}
	// 使用密钥生成claim对应的token串并返回
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}
