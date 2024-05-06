package user_api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/plugins/email"
	"gvb_server/utils/jwt"
	"gvb_server/utils/random"
	"gvb_server/utils/res"
)

type BindEmailRequest struct {
	Email string `json:"email" binding:"required,email" msg:"请输入邮箱"`
	Code  string `json:"code"`
}

// UserBindEmailView 用户绑定邮箱
// @Tags 用户管理
// @Summary 用户绑定邮箱
// @Description 用户绑定邮箱
// @Param data body BindEmailRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/user_bind_email [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserBindEmailView(c *gin.Context) {
	// 绑定参数
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 获取cookie中的载荷信息
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwt.CustomClaims)
	// 启用session
	session := sessions.Default(c)
	// 如果用户没有输入验证码，就向用户邮箱发送验证码
	if cr.Code == "" {
		// 生成四位随机数
		code := random.Code()
		// 将验证码和邮箱保存到session中
		session.Set("valid_code", code)
		session.Set("send_email", cr.Email)
		// 保存session
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("session保存失败", c)
			return
		}
		// 发送邮件
		err = email.NewCode().Send(cr.Email, code)
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("发送邮件失败", c)
			return
		}
		res.OKWithMsg("发送邮件成功", c)
		return
	}
	// 用户输入了验证码就判断用户输入的验证码和session中保存的是否一致
	code := session.Get("valid_code")
	email := session.Get("send_email")
	if email != cr.Email {
		res.FailWithMsg("邮箱输入错误", c)
		return
	}
	if code != cr.Code {
		res.FailWithMsg("验证码输入错误", c)
		return
	}
	var userModel models.UserModel
	count := global.DB.Take(&userModel, claim.UserID).RowsAffected
	if count == 0 {
		res.FailWithMsg("用户不存在，无法为用户绑定邮箱", c)
		return
	}
	err = global.DB.Model(&userModel).Updates(map[string]any{
		"email": cr.Email,
	}).Error
	if err != nil {
		res.FailWithMsg("修改用户失败", c)
		return
	}
	res.OKWithMsg("修改用户成功", c)
}
