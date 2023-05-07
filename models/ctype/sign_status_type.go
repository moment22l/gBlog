package ctype

import (
	"encoding/json"
)

type SignStatus int

const (
	SignQQ    SignStatus = 1 // QQ
	SignGitee SignStatus = 2 // gitee
	SignEmail SignStatus = 3 // 邮箱
)

func (r SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r SignStatus) String() string {
	var str string
	switch r {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "gitee"
	case SignEmail:
		str = "邮箱"
	default:
		str = "其他"
	}
	return str
}
