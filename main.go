package main

import (
	"gBlog/core"
	"gBlog/global"
)

func main() {
	// 初始化配置
	core.InitConfig()
	// 初始化gorm
	global.DB = core.InitGorm()
}
