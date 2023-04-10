package common

import (
	"errors"
	"gBlog/global"
	"gBlog/models"

	"gorm.io/gorm"
)

// Option 分页选项
type Option struct {
	models.PageInfo
	Debug bool
}

// ComList 分页
func ComList[T any](model T, option Option) (list []T, err error) {
	db := global.DB
	// 设置日志模式是否为debug
	if option.Debug {
		db = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	// 设置排序模式
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照创建的时间排序 顺序(从晚到早)
	}

	total := db.Where(model).Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if int64(offset) >= total {
		err = errors.New("页码过大")
		return
	}
	err = db.Where(model).Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return
}
