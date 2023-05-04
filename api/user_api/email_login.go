package user_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/jwts"
	"gBlog/utils/pwd"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

// EmailLoginRequest 邮箱登录所需Request
type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// EmailLoginView 邮箱登录
func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 检测用户是否存在
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ?", cr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 校验密码
	hashPwd := userModel.Password
	check := pwd.CheckPwd(hashPwd, cr.Password)
	if !check {
		global.Log.Warn("密码错误")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 生成token并返回给客户端
	token, err := jwts.GenToken(jwts.JwtPayload{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Warn("token生成失败")
		res.FailWithMessage("token生成失败", c)
		return
	}
	res.OkWithData(token, c)
}
