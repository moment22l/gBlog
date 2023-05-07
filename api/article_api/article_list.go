package article_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/service/common"
	"gBlog/utils/error_code"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

// ArticleListView 查看文章列表
func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}
	// 加载文章时需要同时加载文章的标签, 评论, 作者以及封面
	list, err := common.ComList(models.ArticleModel{},
		global.DB.Preload("TagModels").Preload("CommentModels").Preload("UserModel").Preload("Banner"), common.Option{
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
