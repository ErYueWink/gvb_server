package user_ser

import (
	"errors"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils"
	"gvb_server/utils/pwd"
)

func (UserService) CreateUser(username, password, nickname, email string, role ctype.Role, ip string) error {
	// 查询用户是否已存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", username).Error
	if err == nil {
		global.Log.Error("用户已存在")
		return errors.New("用户已存在")
	}

	// 对密码进行hash
	defPwd := pwd.HashPwd(password)
	avatar := "https://tse1-mm.cn.bing.net/th/id/OIP-C.3gI8DGT2YcMJFYYTG4u0MQAAAA?pid=ImgDet&w=400&h=400&rs=1"
	addr := utils.GetAddr(ip)
	err = global.DB.Create(&models.UserModel{
		UserName:   username,
		NickName:   nickname,
		Password:   defPwd,
		Email:      email,
		Avatar:     avatar,
		Role:       role,
		IP:         ip,
		Addr:       addr,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
