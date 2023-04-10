package advert_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/service/common"
	"gBlog/utils/error_code"
	"gBlog/utils/res"
	"strings"

	"github.com/gin-gonic/gin"
)

// AdvertListView 广告列表
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}
	// 判断 Referer 是否包含admin, 有则返回list的全部信息, 无则只返回IsShow字段
	referer := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		isShow = false
	}
	list, err := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
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
