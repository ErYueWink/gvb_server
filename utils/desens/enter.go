package desens

import (
	"gvb_server/global"
	"strings"
)

// DesensitizationEmail 邮箱脱敏
func DesensitizationEmail(email string) string {
	emailArr := strings.Split(email, "@")
	if len(emailArr) != 2 {
		global.Log.Error("邮箱格式有误")
		return ""
	}
	return emailArr[0][:1] + "*****@" + emailArr[1]
}

// DesensitizationPhone 手机号脱敏
func DesensitizationPhone(phone string) string {
	if len(phone) != 11 {
		global.Log.Error("手机号格式有误")
		return ""
	}
	return phone[:3] + "****" + phone[7:]
}
