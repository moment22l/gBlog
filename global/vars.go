package global

import (
	"gBlog/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Conf *config.Config
	DB   *gorm.DB
	Log  *logrus.Logger
)
