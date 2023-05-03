package routers

import "gBlog/api"

// MenuRouter 菜单路由
func (r RouterGroup) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	r.POST("menu/create", menuApi.MenuCreateView)
	r.GET("menu/list", menuApi.MenuListView)
	r.GET("menu/listName", menuApi.MenuNameListView)
	r.PUT("menu/:id", menuApi.MenuUpdateView)
	r.DELETE("menu/remove", menuApi.MenuRemoveView)
	r.GET("menu/:id", menuApi.MenuDetailView)
}
