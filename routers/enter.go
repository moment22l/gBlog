package routers

import (
	"gBlog/global"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

// RouterGroup 路由组
type RouterGroup struct {
	*gin.RouterGroup
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	gin.SetMode(global.Conf.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 路由分组
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 路由分层
	routerGroupApp.SettingsRouter() // 系统配置
	routerGroupApp.ImagesRouter()   // 图片
	routerGroupApp.AdvertRouter()   // 广告
	routerGroupApp.MenuRouter()     // 菜单
	routerGroupApp.UserRouter()     // 用户
	routerGroupApp.TagRouter()      // 标签
	return router
}
