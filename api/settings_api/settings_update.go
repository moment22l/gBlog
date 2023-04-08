package settings_api

import (
	"gBlog/config"
	"gBlog/core"
	"gBlog/global"
	"gBlog/utils/error_code"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

// SettingsInfoUpdateView 修改配置文件视图
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	// 绑定Uri
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}
	// 按照uri中的name去修改对应配置信息
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			global.Log.Error(err)
			res.FailWithCode(error_code.ArgumentError, c)
			return
		}
		global.Conf.SiteInfo = info
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
	err = core.ModifyConf()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.Ok(c)
}
