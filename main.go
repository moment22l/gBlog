package main

import (
	"gBlog/core"
	"gBlog/global"
	"gBlog/routers"
)

func main() {
	// 初始化配置
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	// 初始化gorm, 连接数据库
	global.DB = core.InitGorm()
	// 初始化routers
	router := routers.InitRouter()
	addr := global.Conf.System.Addr()
	global.Log.Infof("server run on %s", addr)
	router.Run(addr)
}
