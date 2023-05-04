package routers

import (
	"gBlog/api"
	"gBlog/middleware"
)

func (r RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	r.POST("email_login", userApi.EmailLoginView)
	r.GET("user/list", middleware.JwtAuth(), userApi.UserListView)
	r.PUT("user/update_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	r.PUT("user/update_pwd", middleware.JwtAuth(), userApi.UserUpdatePwdView)
	r.POST("user/logout", middleware.JwtAuth(), userApi.UserLogoutView)
	r.DELETE("user/remove", middleware.JwtAdmin(), userApi.UserRemoveView)
}
