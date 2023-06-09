package advert_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/res"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// AdvertUpdateView 更新广告内容
// @Tags 广告管理
// @Summary 更新广告内容
// @param data body AdvertRequest true "更新后的广告信息"
// @Router /api/advert/:id [PUT]
// @Produce json
// @Success 200 {object} res.Response{msg=string}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 判断是否存在该广告
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		global.Log.Errorf("标题为\"%s\"的广告不存在", cr.Title)
		res.FailWithMessage("该广告不存在", c)
		return
	}

	// 更新内容
	m := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(m).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败", c)
		return
	}
	res.OkWithMessage("修改广告成功", c)
}
