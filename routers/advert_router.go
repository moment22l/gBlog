package routers

import (
	"gBlog/api"
	"gBlog/middleware"
)

// AdvertRouter 广告路由
func (r RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdvertApi
	r.POST("advert/create", middleware.JwtAdmin(), app.AdvertCreateView)
	r.GET("advert/list", app.AdvertListView)
	r.PUT("advert/:id", middleware.JwtAdmin(), app.AdvertUpdateView)
	r.DELETE("advert/remove", middleware.JwtAdmin(), app.AdvertRemoveView)
}
