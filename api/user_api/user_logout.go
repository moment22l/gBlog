package user_api

import (
	"fmt"
	"gBlog/global"
	"gBlog/utils/jwts"
	"gBlog/utils/res"
	"time"

	"github.com/gin-gonic/gin"
)

// UserLogoutView 用户登出
func (UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 需要计算距离现在的过期时间并将其加入到redis中
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	err := global.Redis.Set(fmt.Sprintf("logout_%s", c.Request.Header.Get("token")), "", diff).Err()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)
}
