package images_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"未选择图片ID"`
	Name string `json:"name" binding:"required" msg:"未输入图片名称"`
}

// ImagesUpdateView 修改图片名称
// @Tags 图片管理
// @Summary 修改图片名称
// @param data body ImageUpdateRequest true "修改名称所需参数"
// @Router /api/images/update [PUT]
// @Produce json
// @Success 200 {object} res.Response{msg=string}
func (ImagesApi) ImagesUpdateView(c *gin.Context) {
	// 获取参数
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 找到对应id的图片
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("图片不存在", c)
		return
	}
	// 更改对应图片的name字段
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("图片名称修改成功", c)
}
