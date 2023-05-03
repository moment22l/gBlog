package flag

import (
	"fmt"
	"gBlog/global"
	"gBlog/models"
	"gBlog/models/ctype"
	"gBlog/utils/pwd"
)

func CreateUser(permissions string) {
	// 用户名 昵称 密码 确认密码 邮箱
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("请输入用户名: ")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称: ")
	fmt.Scan(&nickName)
	fmt.Printf("请输入密码: ")
	fmt.Scan(&password)
	fmt.Printf("请输入确认密码: ")
	fmt.Scan(&rePassword)
	fmt.Printf("请输入邮箱: ")
	fmt.Scan(&email)

	// 判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		global.Log.Error("用户名已存在, 请重新输入")
		return
	}
	// 校验两次密码是否一致
	if password != rePassword {
		global.Log.Error("两次密码不一致, 请重新输入")
		return
	}
	// 对密码进行hash
	hashPwd := pwd.HashPwd(password)

	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	// 头像
	// 1. 默认头像
	avatar := "/uploads/avatar/default.png"
	// 2. 随机选择头像
	// 入库
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		Role:       role,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户 %s 创建成功", userName)
}
