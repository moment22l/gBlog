package tag_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标签名" structs:"title"` // 标题
}

// TagCreateView 创建标签
func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 判断是否重复
	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		global.Log.Warnf("名为\"%s\"的标签已存在", cr.Title)
		res.FailWithMessage("该标签已存在", c)
		return
	}
	// 创建标签
	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加标签失败", c)
		return
	}
	res.OkWithMessage("添加标签成功", c)
}
