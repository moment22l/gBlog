package models

import "time"

// User2Collects 自定义第三张表 用户何时收藏何文章
type User2Collects struct {
	UserID       uint         `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"foreignKey:UserID"`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
