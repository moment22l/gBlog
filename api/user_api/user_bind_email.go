package user_api

import (
	"fmt"
	"gBlog/global"
	"gBlog/models"
	"gBlog/plugins/email"
	"gBlog/utils/jwts"
	"gBlog/utils/random"
	"gBlog/utils/res"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// BindEmailRequest 需要的参数
type BindEmailRequest struct {
	Email string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code  *string `json:"code"`
}

// UserBindEmailView 绑定邮箱
func (UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 获取json
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 用户绑定邮箱, 第一次输入邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		global.Log.Error("用户不存在")
		res.FailWithMessage("用户不存在", c)
		return
	}
	session := sessions.Default(c)
	// 后台给邮箱发验证码
	if cr.Code == nil {
		// 生成4位验证码, 将生成的验证码存入session
		code := random.Code(4)
		session.Set("valid_code", code)
		session.Set("email", cr.Email)
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("session错误", c)
			return
		}
		// 发验证码
		err = email.NewCode().Send(cr.Email, fmt.Sprintf("您的验证码是: %s", code))
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("发送验证码失败", c)
			return
		}
		fmt.Println(code)
		res.OkWithMessage("验证码已发送", c)
		return
	}
	code := session.Get("valid_code")
	// 第二次, 用户输入邮箱+验证码
	if *cr.Code != code {
		global.Log.Error("验证码错误")
		res.FailWithMessage("验证码错误", c)
		return
	}
	// 判断第二次request的邮箱是否与第一次一样
	emailAddr := session.Get("email")
	if emailAddr != cr.Email {
		global.Log.Error("前后邮箱不一致")
		res.FailWithMessage("邮箱已修改, 请重新获取验证码", c)
		return
	}
	// 将邮箱更新到当前用户的数据库中
	err = global.DB.Model(&user).Update("email", cr.Email).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("邮箱绑定失败", c)
		return
	}
	res.OkWithMessage("邮箱绑定成功", c)
}
