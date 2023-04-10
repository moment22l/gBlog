package api

import (
	"gBlog/api/advert_api"
	"gBlog/api/images_api"
	"gBlog/api/settings_api"
)

// Group api组
type Group struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
}

// ApiGroupApp api组应用
var ApiGroupApp = new(Group)
