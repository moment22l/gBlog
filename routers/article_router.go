package routers

import (
	"gBlog/api"
	"gBlog/middleware"
)

// ArticleRouter 文章路由
func (r RouterGroup) ArticleRouter() {
	articleApi := api.ApiGroupApp.ArticleApi
	r.POST("article/create", middleware.JwtAuth(), articleApi.ArticleCreateView)
	r.GET("article/list", articleApi.ArticleListView)
	r.DELETE("article/remove", middleware.JwtAuth(), articleApi.ArticleRemoveView)
}
