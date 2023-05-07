package article_api

import (
	"fmt"
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils/error_code"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ArticleRemoveView 删除文章
func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr models.RemoveList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}
	// 获取需要删除的文章列表
	var articleList []models.ArticleModel
	count := global.DB.Find(&articleList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("所有文章均不存在", c)
		return
	}
	// 创建事务删除所有文章
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		for _, article := range articleList {
			// 删除文章与标签的引用
			err = global.DB.Model(&article).Association("TagModels").Clear()
			if err != nil {
				return err
			}
			// 删除文章
			global.DB.Delete(&article)
		}
		return nil
	})
	if err != nil {
		global.Log.Error("删除文章失败")
		res.FailWithMessage("删除文章失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 条文章", count), c)
}
