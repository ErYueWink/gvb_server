package redis_ser

import (
	"gvb_server/global"
	"gvb_server/utils"
	"time"
)

const prefix = "logout_"

// Logout 用户退出登录
func Logout(token string, diff time.Duration) error {
	return global.Redis.Set(prefix+token, "", diff).Err()
}

// CheckLogout 判断用户是否退出登录
func CheckLogout(token string) bool {
	// 获取redis中所有的key
	keys := global.Redis.Keys(prefix + "*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false
}
