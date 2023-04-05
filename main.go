package main

import (
	"fmt"
	"gBlog/core"
	"gBlog/flag"
	"gBlog/global"
	"gBlog/routers"
	"gBlog/utils/error_code"
)

func main() {
	// 初始化配置
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	// 初始化gorm, 连接数据库
	global.DB = core.InitGorm()
	// 初始化错误码信息
	errorMap, err := error_code.InitErrorCode()
	if err != nil {
		global.Log.Warn(err)
	} else {
		global.ErrorMap = errorMap
	}
	// 命令行参数绑定
	option := flag.Parse()
	fmt.Println(option)
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	// 初始化routers
	router := routers.InitRouter()
	addr := global.Conf.System.Addr()
	global.Log.Infof("server run on %s", addr)
	err = router.Run(addr)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
}
