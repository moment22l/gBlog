package core

import (
	"gBlog/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitGorm 初始化gorm
func InitGorm() *gorm.DB {
	// 判断host是否不为空
	if global.Conf.Mysql.Host == "" {
		global.Log.Error("mysql's host is missing")
		return nil
	}
	// 设置日志模式
	var l logger.Interface
	if global.Conf.System.Env == "debug" {
		l = logger.Default.LogMode(logger.Info)
	} else {
		l = logger.Default.LogMode(logger.Warn)
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info)
	// 拿到mysql的dsn
	dsn := global.Conf.Mysql.Dsn()
	// 连接mysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		global.Log.Error("open mysql fail")
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.Conf.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.Conf.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
