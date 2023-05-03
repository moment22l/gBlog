package menu_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")

	// 先之前的banner清空
	var menuModel models.MenuModel
	err = global.DB.Preload("Banners").Take(&menuModel, id).Error
	if err != nil {
		global.Log.Error("菜单id不存在")
		res.FailWithMessage("菜单不存在", c)
		return
	}
	err = global.DB.Model(&menuModel).Association("Banners").Clear()
	// 如果选择了banner, 那就添加
	if len(cr.ImageSortList) > 0 {
		var menuBannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			menuBannerList = append(menuBannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&menuBannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("新菜单图片关联失败", c)
			return
		}
	}
	// 更新除了banner以外的其他数据
	m := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(m).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}
	res.OkWithMessage("修改菜单成功", c)
}
