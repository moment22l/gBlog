package models

import (
	"gBlog/models/ctype"
	"os"

	"gorm.io/gorm"
)

// BannerModel 图片
type BannerModel struct {
	Model
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片hash值
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片存储类型
}

// BeforeDelete hook(钩子)函数, 在调用数据库删除前会先调用该函数
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		// 本地图片删除时, 需要删本地的存储
		err = os.Remove(b.Path)
		if err != nil {
			return
		}
	}
	return
}
