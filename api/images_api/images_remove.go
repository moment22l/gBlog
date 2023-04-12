package images_api

import (
	"fmt"
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/error_code"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

// TODO: 删除七牛的逻辑

// ImagesRemoveView 删除图片
// @Tags 图片管理
// @Summary 删除图片
// @param data body models.RemoveList true "需要删除的图片的id列表"
// @Router /api/images/remove [DELETE]
// @Produce json
// @Success 200 {object} res.Response{msg=string}
func (ImagesApi) ImagesRemoveView(c *gin.Context) {
	var cr models.RemoveList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("所有文件均不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 张图片", count), c)
}
