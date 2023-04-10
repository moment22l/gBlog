package routers

import "gBlog/api"

func (r RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertApi
	r.POST("advert/create", advertApi.AdvertCreateView)
}
