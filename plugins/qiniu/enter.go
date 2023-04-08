package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gBlog/config"
	"gBlog/global"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// getToken 获取文件上传的Token
func getToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	// 鉴权
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// getCfg 获取文件上传配置
func getCfg(q config.QiNiu) storage.Config {
	cfg := storage.Config{}
	// 空间对应机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否需要cdn加速
	cfg.UseCdnDomains = false
	return cfg
}

// UploadImage 上传图片
// data: 文件数组
// prefix: 前缀
func UploadImage(data []byte, name string, prefix string) (filepath string, err error) {
	q := global.Conf.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("未配置AccessKey或SecretKey")
	}
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("文件超过设定大小")
	}
	upToken := getToken(q)
	cfg := getCfg(q)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	// 获取当前时间
	now := time.Now().Format("20230102030405")
	key := fmt.Sprintf("%s/%s_%s", prefix, now, name)

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil
}
