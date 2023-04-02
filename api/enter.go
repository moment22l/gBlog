package api

import "gBlog/api/settings_api"

type Group struct {
	SettingsApi settings_api.SettingsApi
}

var ApiGroupApp = new(Group)
