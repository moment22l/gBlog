package image_svc

import (
	"fmt"
	"gBlog/global"
	"gBlog/models"
	"gBlog/models/ctype"
	"gBlog/plugins/qiniu"
	"gBlog/utils"
	"io"
	"path"
	"strings"

	"mime/multipart"
)

// FileUploadResponse 对上传文件的响应
type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

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

// ImageUploadService 文件上传的方法
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	// 拿到完整的filePath
	filePath := path.Join(global.Conf.Upload.Path, fileName)
	res.FileName = filePath
	res.IsSuccess = false

	// 判断图片后缀是否在白名单中
	ext := strings.ToLower(path.Ext(fileName))
	if !utils.InStringList(ext, whiteImageList) {
		res.Msg = "非法文件"
		return
	}

	// 读取文件内容并计算图片内容的md5
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
		return
	}
	byteData, _ := io.ReadAll(fileObj)
	fileHash := utils.MD5(byteData)

	// 判断图片大小是否超过设定值
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Conf.Upload.Size) {
		res.Msg = fmt.Sprintf("图片超过设定大小 当前大小: %.2fMB 设定大小: %dMB", size, global.Conf.Upload.Size)
		return
	}

	// 判断图片是否存在于数据库当中
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", fileHash).Error
	if err == nil {
		global.Log.Error("所上传的图片已存在")
		res.FileName = bannerModel.Path
		res.Msg = "图片已存在"
		return
	}

	fileType := ctype.Local
	res.Msg = "图片上传本地成功"
	// 检查是否将保存图片到七牛
	if global.Conf.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Conf.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "上传七牛成功"
		fileType = ctype.QiNiu
	}
	res.IsSuccess = true
	// 图片入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      fileHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
