package models

import "gBlog/models/ctype"

type ArticleModel struct {
	Model
	Title         string         `gorm:"size:32" json:"title"`                    // 文章标题
	Abstract      string         `json:"abstract"`                                // 简介
	Content       string         `json:"content"`                                 // 内容
	LookCount     int            `json:"look_count"`                              // 浏览量
	CommentCount  int            `json:"comment_count"`                           // 评论量
	DiggCount     int            `json:"digg_count"`                              // 点赞量
	CollectsCount int            `json:"collects_count"`                          // 收藏量
	TagModels     []TagModel     `gorm:"many2many:article_tag" json:"tag_models"` // 文章标签
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`           // 文章的评论列表
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`              // 文章作者
	UserID        uint           `json:"user_id"`                                 // 用户id
	Link          string         `json:"link"`                                    // 原文链接
	Banner        BannerModel    `gorm:"foreignKey:BannerID" json:"-"`            // 文章封面
	BannerID      uint           `json:"banner_id"`                               // 封面ID
	NickName      string         `gorm:"size:36" json:"nick_name"`                // 发布文章的用户昵称
	BannerPath    string         `json:"banner_path"`                             // 文章封面的路径
	Tags          ctype.Array    `gorm:"type:string;size:64" json:"tags"`         // 文章标签
}
