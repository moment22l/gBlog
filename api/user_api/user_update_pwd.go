package user_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/jwts"
	"gBlog/utils/pwd"
	"gBlog/utils/res"
	"github.com/gin-gonic/gin"
)

type UpdatePwdRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"`
	NewPwd string `json:"new_pwd" binding:"required" msg:"请输入新密码"`
	RePwd  string `json:"re_pwd" binding:"required" msg:"请输入确认密码"`
}

func (UserApi) UserUpdatePwdView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 拿到token对应的user
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("用户不存在", c)
		return
	}
	// 参数绑定
	var cr UpdatePwdRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 判断旧密码是否输入正确
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		global.Log.Error("旧密码输入错误")
		res.FailWithMessage("旧密码输入错误", c)
		return
	}
	// 判断新密码与确认密码是否一致
	if cr.NewPwd != cr.RePwd {
		global.Log.Error("新密码与确认密码不一致")
		res.FailWithMessage("新密码与确认密码不一致", c)
		return
	}
	// 更新密码
	hashPwd := pwd.HashPwd(cr.NewPwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改密码失败", c)
		return
	}
	res.OkWithMessage("修改密码成功", c)
}
