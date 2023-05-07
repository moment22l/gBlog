package ctype

import (
	"encoding/json"
)

type Role int

const (
	PermissionAdmin       Role = 1 // 管理员
	PermissionUser        Role = 2 // 用户
	PermissionVisitor     Role = 3 // 游客
	PermissionDisableUser Role = 4 // 黑名单用户
)

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Role) String() string {
	var str string
	switch r {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisableUser:
		str = "被禁用户"
	default:
		str = "其他"
	}
	return str
}
