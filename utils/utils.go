package utils

// arg:key 图片后缀名 list:图片后缀名白名单列表
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}
