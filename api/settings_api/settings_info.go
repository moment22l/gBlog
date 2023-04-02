package settings_api

import (
	"gBlog/utils/common"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	common.Ok(c)
}
