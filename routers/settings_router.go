package routers

import (
	"gBlog/api"
)

// SettingsRouter 配置相关路由
func (r RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	r.GET("settings/:name", settingsApi.SettingsInfoView)
	r.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
}
