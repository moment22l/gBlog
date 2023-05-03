package menu_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenusResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	// 先查菜单
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	// 查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	var menus []MenusResponse
	for _, model := range menuList {
		banners := make([]Banner, 0)
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenusResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}
	res.OkWithData(menus, c)
}
