package user_api

import (
	"fmt"
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/error_code"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserRemoveView 删除用户
func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}

	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("所有用户均不存在", c)
		return
	}

	// 事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO：删除文章
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("删除用户失败", c)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除用户失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("成功删除 %d 个用户", count), c)
}
