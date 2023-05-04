package user_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/models/ctype"
	"gBlog/utils/res"
	"github.com/gin-gonic/gin"
)

// UserRole 权限变更request
type UserRole struct {
	UserID   uint       `json:"user_id" binding:"required" msg:"请输入用户id"`
	NickName string     `json:"nick_name"`
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
}

// UserUpdateRoleView 用户权限变更
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 判断用户是否存在
	var user models.UserModel
	err = global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("用户不存在", c)
		return
	}
	// 更新数据
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改权限失败", c)
		return
	}
	res.OkWithMessage("修改权限成功", c)
}
