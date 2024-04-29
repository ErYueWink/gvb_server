package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"gvb_server/global"
)

// HashPwd 对密码进行hash加密
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		global.Log.Error("密码hash失败")
		return ""
	}
	return string(hash)
}

// CheckPwd 校验密码
func CheckPwd(hashPwd, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		global.Log.Error(err)
		return false
	}
	return true
}
