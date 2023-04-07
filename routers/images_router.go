package routers

import "gBlog/api"

// ImagesRouter 图片路由
func (r RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	r.POST("images/upload", imagesApi.ImagesUploadView)
}
