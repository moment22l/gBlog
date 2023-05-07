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
	r.POST("user/login", userApi.UserLoginView)                                  // 用户登录
	r.GET("user/list", middleware.JwtAuth(), userApi.UserListView)               // 查看用户列表
	r.PUT("user/update_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView) // 用户权限更新
	r.PUT("user/update_pwd", middleware.JwtAuth(), userApi.UserUpdatePwdView)    // 用户修改密码
	r.POST("user/logout", middleware.JwtAuth(), userApi.UserLogoutView)          // 用户登出
	r.DELETE("user/remove", middleware.JwtAdmin(), userApi.UserRemoveView)       // 删除用户
	r.POST("user/bind_email", middleware.JwtAuth(), userApi.UserBindEmailView)   // 绑定邮箱
	r.POST("user/create", userApi.UserCreateView)                                // 创建用户
}
