package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// JwtPayload jwt中的payload数据
type JwtPayload struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role"` // 权限
	UserID   uint   `json:"user_id"`
}

// CustomClaims  jwt信息
type CustomClaims struct {
	JwtPayload
	jwt.StandardClaims
}

var MySecret []byte
