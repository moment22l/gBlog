package menu_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/res"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	// 先查菜单
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		global.Log.Error("菜单不存在")
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// 查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
	var menu MenusResponse
	banners := make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menu = MenusResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menu, c)
}
