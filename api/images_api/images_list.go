package images_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/service/common"
	"gBlog/utils/error_code"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

// ImagesListView 图片列表
func (ImagesApi) ImagesListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}
	list, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithList(list, int64(len(list)), c)

}
