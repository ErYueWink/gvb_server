package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
	"gvb_server/utils/res"
)

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// EmailLoginView 用户名或者邮箱登录
// @Tags 用户管理
// @Summary 邮箱登录
// @Description 邮箱登录，返回token，用户信息需要从token中解码
// @Param data body EmailLoginRequest  true  "查询参数"
// @Router /api/email_login [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) EmailLoginView(c *gin.Context) {
	var cr UserLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	// 根据前端传来的用户名进行查询
	var userModel models.UserModel
	count := global.DB.Take(&userModel, "user_name = ? or email = ?", cr.Username, cr.Username).RowsAffected
	if count == 0 {
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	// 校验密码
	flag := pwd.CheckPwd(userModel.Password, cr.Password)
	if !flag {
		res.FailWithMsg("用户名或密码失败", c)
		return
	}
	// 登录成功生成token
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		UserID:   userModel.ID,
		Username: userModel.UserName,
		Role:     int(userModel.Role),
		NickName: userModel.NickName,
		Avatar:   userModel.Avatar,
	})
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMsg(err.Error(), c)
		return
	}
	// 获取登录IP和登录地址
	ip, addr := utils.GetAddrByGin(c)
	// 将token保存到请求头
	c.Request.Header.Set("token", token)
	// 保存登录日志
	err = global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,         // 用户ID
		NickName:  userModel.NickName,   // 用户昵称
		Device:    "",                   // 登录设备
		IP:        ip,                   // 登录IP
		Addr:      addr,                 // 登录地址
		Token:     token,                // token
		LoginType: userModel.SignStatus, // 登录类型
	}).Error
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMsg("创建登录日志失败", c)
		return
	}
	res.OKWithData(token, c)
}
