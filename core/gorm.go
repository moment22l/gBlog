package core

import (
	"gBlog/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Conf.Mysql.Host == "" {
		global.Log.Error("mysql's host is missing")
		return nil
	}
	var l logger.Interface
	if global.Conf.System.Env == "debug" {
		l = logger.Default.LogMode(logger.Info)
	} else {
		l = logger.Default.LogMode(logger.Warn)
	}
	dsn := global.Conf.Mysql.Dsn()
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
