package user_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/models/ctype"
	"gBlog/service/common"
	"gBlog/utils/desensitization"
	"gBlog/utils/error_code"
	"gBlog/utils/jwts"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

// UserListView 用户列表
func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}
	// list分页查询用户列表
	var users []models.UserModel
	list, err := common.ComList(models.UserModel{}, common.Option{
		PageInfo: cr,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	for _, user := range list {
		if claims.Role != int(ctype.PermissionAdmin) {
			// 管理员
			user.UserName = ""
		}
		// 信息脱敏
		user.Tel = desensitization.Tel(user.Tel)
		user.Email = desensitization.Email(user.Email)
		users = append(users, user)
	}

	res.OkWithList(users, int64(len(users)), c)
}
