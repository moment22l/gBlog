package routers

import (
	"gBlog/api"
	"gBlog/middleware"
)

// TagRouter 标签路由
func (r RouterGroup) TagRouter() {
	tagApi := api.ApiGroupApp.TagApi
	r.POST("tag/create", middleware.JwtAdmin(), tagApi.TagCreateView)
	r.GET("tag/list", middleware.JwtAdmin(), tagApi.TagListView)
	r.PUT("tag/:id", middleware.JwtAdmin(), tagApi.TagUpdateView)
	r.DELETE("tag/remove", middleware.JwtAdmin(), tagApi.TagRemoveView)
}
