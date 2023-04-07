package config

// Upload 上传图片配置
type Upload struct {
	Size      int    `json:"size" yaml:"size"` // 图片上传的大小
	Path      string `json:"path" yaml:"path"` // 图片上传的目录
	StorePath string `json:"store_path" yaml:"store_path"`
}
