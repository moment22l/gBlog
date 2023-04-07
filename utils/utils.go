package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 md5加密
func MD5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// InStringList 判断key是否存在于list当中
func InStringList(key string, list []string) bool {
	for _, s := range list {
		if s == key {
			return true
		}
	}
	return false
}
