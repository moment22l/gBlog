package main

import (
	"gBlog/core"
	"gBlog/global"
)

func main() {
	// 初始化配置
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("xxx")
	global.Log.Error("xxx")
	global.Log.Info("xxx")
	// 初始化gorm
	global.DB = core.InitGorm()
}
