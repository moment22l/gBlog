package settings_api

import (
	"gBlog/global"
	"gBlog/utils/error_code"
	"gBlog/utils/res"

	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 根据uri中的name显示对应的配置信息
// @Tags 系统配置管理
// @Summary 查看配置信息
// @Router /api/settings/:name [GET]
// @Produce json
// @Success 200 {object} res.Response{data=config.SiteInfo, msg=string}
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(error_code.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Conf.SiteInfo, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
}
