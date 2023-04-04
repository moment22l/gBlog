package models

// AdvertModel 广告表
type AdvertModel struct {
	Model
	Title  string `gorm:"size:32" json:"title"` // 标题
	Href   string `json:"href"`                 // 超链接
	Images string `json:"images"`               // 图片
	IsShow bool   `json:"is_show"`              // 是否展示
}
