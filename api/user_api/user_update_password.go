package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
	"gvb_server/utils/res"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"`
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`
}

// UserUpdatePasswordView 修改登录人密码
// @Tags 用户管理
// @Summary 修改登录人的密码
// @Description 修改登录人的密码
// @Param data body UpdatePasswordRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/user_password [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	var cr UpdatePasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 获取Cookie中的载荷信息
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	// 根据用户ID查询用户是否存在
	var userModel models.UserModel
	count := global.DB.Take(&userModel, claims.UserID).RowsAffected
	if count == 0 {
		res.FailWithMsg("用户不存在", c)
		return
	}
	// 校验旧密码和数据库中保存的密码是否一致
	flag := pwd.CheckPwd(userModel.Password, cr.OldPwd)
	if !flag {
		res.FailWithMsg("原密码输入错误", c)
		return
	}
	// 密码输入正确的情况下，修改密码
	err = global.DB.Model(&userModel).Update("password", pwd.HashPwd(cr.Pwd)).Error
	if err != nil {
		res.FailWithMsg("修改密码失败", c)
		return
	}
	res.OKWithMsg("修改密码成功", c)

}
