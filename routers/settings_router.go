package routers

import (
	"gBlog/api"
	"gBlog/middleware"
)

// SettingsRouter 配置相关路由
func (r RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	r.GET("settings/:name", middleware.JwtAdmin(), settingsApi.SettingsInfoView)
	r.PUT("settings/:name", middleware.JwtAdmin(), settingsApi.SettingsInfoUpdateView)
}
