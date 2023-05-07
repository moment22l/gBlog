package article_api

import (
	"gBlog/global"
	"gBlog/models"
	"gBlog/models/ctype"
	"gBlog/utils/jwts"
	"gBlog/utils/res"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ArticleCreateReq 文章创建请求
type ArticleCreateReq struct {
	Title    string      `json:"title" binding:"required" msg:"请输入文章标题"`   // 文章标题
	Abstract string      `json:"abstract"`                                 // 简介
	Content  string      `json:"content" binding:"required" msg:"请输入文章内容"` // 内容
	Tags     ctype.Array `json:"tag"`                                      // 标签
	BannerID uint        `json:"banner_id"`                                // 封面ID
}

// ArticleCreateView 创建文章
func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ArticleCreateReq
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// token信息
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 获取用户
	var user models.UserModel
	global.DB.Take(&user, claims.UserID)
	// 标签列表
	var tagList []models.TagModel
	var tags ctype.Array
	for _, title := range cr.Tags {
		var tag models.TagModel
		err = global.DB.Take(&tag, "title = ?", title).Error
		if err != nil {
			continue
		}
		tagList = append(tagList, tag)
		tags = append(tags, tag.Title)
	}
	// 封面是否存在
	var banner models.BannerModel
	err = global.DB.Take(&banner, cr.BannerID).Error
	if err != nil {
		global.Log.Error("封面图片不存在")
		res.FailWithMessage("封面图片不存在", c)
		return
	}
	// 创建文章
	// TODO: 生成文章链接URL
	article := models.ArticleModel{
		Title:      cr.Title,
		Abstract:   cr.Abstract,
		Content:    cr.Content,
		TagModels:  tagList,
		UserModel:  user,
		Link:       "",
		Banner:     banner,
		NickName:   claims.NickName,
		BannerPath: banner.Path,
		Tags:       tags,
	}
	// 开启事务, 完成创建文章并将该文章加入到对应用户的文章列表中
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Create(&article).Error
		if err != nil {
			return err
		}
		err = global.DB.Model(&user).Association("ArticleModels").Append(&article)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error("文章创建失败")
		res.FailWithMessage("文章创建失败", c)
		return
	}
	res.OkWithMessage("文章创建成功", c)
}
