package ctype

import (
	"encoding/json"
	"strconv"
)

// TODO: 没有起作用

type ImageType int

const (
	Local ImageType = 1 // 本地
	QiNiu ImageType = 2 // 七牛云
)

func (i ImageType) MarshalJson() ([]byte, error) {
	return json.Marshal(strconv.Itoa(int(i)))
}

func (i ImageType) String() string {
	var str string
	switch i {
	case Local:
		str = "本地"
	case QiNiu:
		str = "七牛云"
	default:
		str = "其他"
	}
	return str
}
