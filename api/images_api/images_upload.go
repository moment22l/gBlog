package images_api

import (
	"fmt"
	"gBlog/global"
	"gBlog/models"
	"gBlog/utils"
	"gBlog/utils/common"
	"io"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// whiteImageList 图片白名单
var whiteImageList = []string{
	".jpg",
	".png",
	".jpeg",
	".webp",
	".gif",
	".ico",
	".tiff",
	".svg",
}

// FileUploadResponse 对上传文件的响应
type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// ImagesUploadView 上传多个图片, 返回图片URL
func (ImagesApi) ImagesUploadView(c *gin.Context) {
	// fileHeader, err := c.FormFile("image")
	form, err := c.MultipartForm()
	if err != nil {
		global.Log.Error(err)
		common.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		common.FailWithMessage("文件不存在或未上传文件", c)
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
	var resList []FileUploadResponse
	for _, file := range fileList {
		fileName := file.Filename
		// 判断图片后缀是否在白名单中
		ext := strings.ToLower(path.Ext(fileName))
		if !utils.InStringList(ext, whiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}
		// 计算图片内容的md5
		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
			continue
		}
		byteData, _ := io.ReadAll(fileObj)
		fileHash := utils.MD5(byteData)
		// 判断图片是否存在于数据库当中
		var bannerModel models.BannerModel
		err = global.DB.Take(bannerModel, "hash = ?", fileHash).Error
		if err == nil {
			global.Log.Error()
			resList = append(resList, FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       "图片已存在",
			})
			continue
		}

		filePath := path.Join(basePath, fileName)
		// 判断图片大小是否超过设定值
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Conf.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小过大 当前大小: %.2fMB 设定大小: %dMB", size, global.Conf.Upload.Size),
			})
			continue
		}
		// 保存图片
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "图片上传成功",
		})
		// 图片入库
		global.DB.Create(&models.BannerModel{
			Path: filePath,
			Hash: fileHash,
			Name: fileName,
		})
	}
	common.OkWithData(resList, c)
}
