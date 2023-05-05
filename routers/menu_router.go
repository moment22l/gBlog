package routers

import (
	"gBlog/api"
	"gBlog/middleware"
)

// MenuRouter 菜单路由
func (r RouterGroup) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	r.POST("menu/create", middleware.JwtAdmin(), menuApi.MenuCreateView)
	r.GET("menu/list", middleware.JwtAdmin(), menuApi.MenuListView)
	r.GET("menu/listName", middleware.JwtAdmin(), menuApi.MenuNameListView)
	r.PUT("menu/:id", middleware.JwtAdmin(), menuApi.MenuUpdateView)
	r.DELETE("menu/remove", middleware.JwtAdmin(), menuApi.MenuRemoveView)
	r.GET("menu/:id", middleware.JwtAdmin(), menuApi.MenuDetailView)
}
