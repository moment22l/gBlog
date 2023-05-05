package tag_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

// TagUpdateView 更新标签
func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 判断是否存在该标签
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		global.Log.Errorf("id为\"%s\"的标签不存在", id)
		res.FailWithMessage("该标签不存在", c)
		return
	}

	// 更新内容
	err = global.DB.Model(&tag).Update("title", cr.Title).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改标签失败", c)
		return
	}
	res.OkWithMessage("修改标签成功", c)
}
