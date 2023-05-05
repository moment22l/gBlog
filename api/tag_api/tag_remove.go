package tag_api

import (
	"fmt"
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/error_code"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

// TagRemoveView 批量删除标签
func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}
	// 找到所有存在的标签
	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("所有标签均不存在", c)
		return
	}
	global.DB.Delete(&tagList)
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 条标签", count), c)
}
