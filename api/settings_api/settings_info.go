package settings_api

import (
	"gBlog/global"
	"gBlog/utils/common"
	"gBlog/utils/error_code"

	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 根据uri中的name显示对应的配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error(err)
		common.FailWithCode(error_code.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		common.OkWithData(global.Conf.SiteInfo, c)
	default:
		common.FailWithMessage("没有对应的配置信息", c)
	}
}
