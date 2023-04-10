package routers

import (
	"gBlog/global"

	"github.com/gin-gonic/gin"
)

// RouterGroup 路由组
type RouterGroup struct {
	*gin.RouterGroup
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	gin.SetMode(global.Conf.System.Env)
	router := gin.Default()
	// 路由分组
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 路由分层
	routerGroupApp.SettingsRouter() // 系统配置
	routerGroupApp.ImagesRouter()   // 图片
	routerGroupApp.AdvertRouter()   // 广告
	return router
}
