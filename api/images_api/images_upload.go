package images_api

import (
	"gBlog/global"
	"gBlog/service"
	"gBlog/service/image_svc"
	"gBlog/utils/res"
	"os"

	"github.com/gin-gonic/gin"
)

// ImagesUploadView 上传多个图片, 返回图片URL
func (ImagesApi) ImagesUploadView(c *gin.Context) {
	// fileHeader, err := c.FormFile("image")
	form, err := c.MultipartForm()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("文件不存在或未上传文件", c)
		return
	}
	// 判断路径是否存在
	basePath := global.Conf.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			global.Log.Error(err)
		} else {
			global.Log.Info("文件目录创建成功")
		}
	}

	// 上传图片
	var resList []image_svc.FileUploadResponse
	for _, file := range fileList {
		response := service.ServiceApp.ImageService.ImageUploadService(file)
		// 出错则直接返回错误信息
		if !response.IsSuccess {
			resList = append(resList, response)
			continue
		}
		// 判断是否需要存入本地
		if !global.Conf.QiNiu.Enable {
			err = c.SaveUploadedFile(file, response.FileName)
			if err != nil {
				global.Log.Error(err)
				response.Msg = err.Error()
				response.IsSuccess = false
				resList = append(resList, response)
				continue
			}
		}
		resList = append(resList, response)
	}
	res.OkWithData(resList, c)
}
