package core

import (
	"gBlog/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Conf.Mysql.Host == "" {
		log.Fatalf("无mysql地址")
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
		log.Fatalf("mysql连接失败")
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.Conf.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.Conf.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
