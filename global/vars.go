package global

import (
	"gBlog/config"
	"gBlog/utils/error_code"
	"github.com/go-redis/redis"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Conf     *config.Config      // 整体配置
	DB       *gorm.DB            // 数据库
	Log      *logrus.Logger      // 默认日志
	Redis    *redis.Client       // redis
	ErrorMap error_code.ErrorMap // 错误码表
	MysqlLog logger.Interface    // mysql日志类型
)
