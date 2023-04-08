package models

import "time"

// Model 主模型
type Model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PageInfo 图片分页
type PageInfo struct {
	Page  int    `form:"page"` // 页码
	Key   string `form:"key"`
	Limit int    `form:"limit"` // 一页的数量
	Sort  string `form:"sort"`
}

// RemoveList 图片删除列表
type RemoveList struct {
	IDList []uint `json:"id_list"`
}
