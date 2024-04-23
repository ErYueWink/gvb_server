package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 // 管理员
	PermissionUser        Role = 2 // 普通登录人
	PermissionVisitor     Role = 3 // 游客
	PermissionDisableUser Role = 4 // 被禁用的用户
)

// MarshalJson 转换为json
func (r Role) MarshalJson() ([]byte, error) {
	return json.Marshal(r.String())
}

// String 为枚举添加输出函数，将数字常量转换为字符串
func (r Role) String() string {
	var str string
	switch r {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "普通用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisableUser:
		str = "被禁言的用户"
	default:
		str = "其他"
	}
	return str
}
