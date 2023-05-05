package user_api

import (
	"fmt"
	"gBlog/global"
	"gBlog/models"
	"gBlog/models/ctype"
	"gBlog/utils/pwd"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

// UserCreateReq 创建用户请求
type UserCreateReq struct {
	NickName   string     `json:"nick_name" binding:"required" msg:"请输入昵称"`     // 昵称
	UserName   string     `json:"user_name" binding:"required" msg:"请输入用户名"`    // 用户名
	Password   string     `json:"password" binding:"required" msg:"请输入密码"`      // 密码
	RePassword string     `json:"re_password" binding:"required" msg:"请输入确认密码"` // 密码
	Role       ctype.Role `json:"role" binding:"required" msg:"请输入昵称"`          // 用户角色
	Email      string     `json:"email"`                                        // 邮箱
	Tel        string     `json:"tel"`                                          // 手机号
}

// UserCreateView 创建用户
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateReq
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 检查是否有重复的用户
	var user models.UserModel
	err = global.DB.Take(&user, "user_name = ?", cr.UserName).Error
	if err == nil {
		global.Log.Warn("用户已存在")
		res.FailWithMessage("用户已存在", c)
		return
	}
	// 校验密码是否正确并生成pwd的hash值
	if cr.Password != cr.RePassword {
		global.Log.Warn("确认密码与密码不一致")
		res.FailWithMessage("确认密码与密码不一致", c)
		return
	}
	hashedPwd := pwd.HashPwd(cr.Password)
	// 创建用户
	avatar := "/uploads/avatar/default.png"
	err = global.DB.Create(&models.UserModel{
		NickName:   cr.NickName,
		UserName:   cr.UserName,
		Password:   hashedPwd,
		Email:      cr.Email,
		Tel:        cr.Tel,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		Role:       cr.Role,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error("用户创建失败")
		res.FailWithMessage("用户创建失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("用户 %s 创建成功", cr.UserName), c)
}
