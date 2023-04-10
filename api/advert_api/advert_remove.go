package advert_api

import (
	"fmt"
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/error_code"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

// AdvertRemoveView 删除广告
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}

	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("所有广告均不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 条广告", count), c)
}
