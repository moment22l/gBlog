package api

import (
	"gBlog/api/images_api"
	"gBlog/api/settings_api"
)

// Group api组
type Group struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
}

// ApiGroupApp api组应用
var ApiGroupApp = new(Group)
