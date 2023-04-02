package global

import (
	"gBlog/config"
	"gBlog/utils/error_code"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Conf     *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	ErrorMap error_code.ErrorMap
)
