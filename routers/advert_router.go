package routers

import "gBlog/api"

// AdvertRouter 广告路由
func (r RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdvertApi
	r.POST("advert/create", app.AdvertCreateView)
	r.GET("advert/list", app.AdvertListView)
	r.PUT("advert/:id", app.AdvertUpdateView)
	r.DELETE("advert/remove", app.AdvertRemoveView)
}
