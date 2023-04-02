package global

import (
	"gBlog/config"
	"gorm.io/gorm"
)

var (
	Conf *config.Config
	DB   *gorm.DB
)
