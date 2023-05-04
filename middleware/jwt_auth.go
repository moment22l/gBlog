package middleware

import (
	"fmt"
	"gBlog/global"
	"gBlog/models/ctype"
	"gBlog/utils/jwts"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

// JwtAuth 登录中间件, 供用户调用
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.Request.Header.Get("token")
		if token == "" {
			global.Log.Error("未携带token")
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		// 从token解析claims
		claims, err := jwts.ParseToken(token)
		if err != nil {
			global.Log.Error("token无法解析")
			res.FailWithMessage("token无效", c)
			c.Abort()
			return
		}
		// 判断token是否在redis中
		err = global.Redis.Get(fmt.Sprintf("logout_%s", token)).Err()
		if err == nil {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		// 登录的用户
		c.Set("claims", claims)
	}
}

// JwtAdmin 供超级管理员调用
func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.Request.Header.Get("token")
		if token == "" {
			global.Log.Error("未携带token")
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		// 从token解析claims
		claims, err := jwts.ParseToken(token)
		if err != nil {
			global.Log.Error("token无法解析")
			res.FailWithMessage("token无效", c)
			c.Abort()
			return
		}
		// 判断用户是否为管理员
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("权限不足", c)
			c.Abort()
			return
		}
		// 判断token是否在redis中
		err = global.Redis.Get(fmt.Sprintf("logout_%s", token)).Err()
		if err == nil {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		// 登录的用户
		c.Set("claims", claims)
	}
}
