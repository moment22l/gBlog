package routers

import "gBlog/api"

// ImagesRouter 图片路由
func (r RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	r.POST("images/upload", imagesApi.ImagesUploadView)
	r.GET("images/list", imagesApi.ImagesListView)
	r.DELETE("images/remove", imagesApi.ImagesRemoveView)
	r.PUT("images/update", imagesApi.ImagesUpdateView)
}
