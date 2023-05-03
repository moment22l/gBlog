package menu_api

import (
	"fmt"
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/error_code"
	"gBlog/utils/res"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}

	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("所有菜单均不存在", c)
		return
	}

	// 事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除第三张表
		err = global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("删除关联菜单图片失败", c)
			return err
		}
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("删除菜单失败", c)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("成功删除 %d 条菜单", count), c)
}
