package menu_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/models/ctype"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`
	BannerTime    int         `json:"banner_time" structs:"banner_time"`
	Sort          int         `json:"sort" binding:"required" msg:"请选择菜单图片序号" structs:"sort"`
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 重复判断
	var menu models.MenuModel
	err = global.DB.Take(&menu, "title = ? or path = ?", cr.Title, cr.Path).Error
	if err == nil {
		global.Log.Warnf("标题为\"%s\"或路径为\"%s\"的菜单已存在", cr.Title, cr.Path)
		res.FailWithMessage("同标题或同路径的菜单已存在", c)
		return
	}
	// 创建banner数据入库
	menuModel := &models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OkWithMessage("菜单添加成功", c)
		return
	}
	// 给第三张表入库
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
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}

	res.OkWithMessage("菜单添加成功", c)
}
