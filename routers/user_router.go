package routers

import (
	"gBlog/api"
	"gBlog/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("HuxHS23WiN123nvi1MoPQI980I"))

func (r RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	r.Use(sessions.Sessions("sessionid", store))
	r.POST("email_login", userApi.EmailLoginView)
	r.GET("user/list", middleware.JwtAuth(), userApi.UserListView)
	r.PUT("user/update_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	r.PUT("user/update_pwd", middleware.JwtAuth(), userApi.UserUpdatePwdView)
	r.POST("user/logout", middleware.JwtAuth(), userApi.UserLogoutView)
	r.DELETE("user/remove", middleware.JwtAdmin(), userApi.UserRemoveView)
	r.POST("user/bind_email", middleware.JwtAuth(), userApi.UserBindEmailView)
}
