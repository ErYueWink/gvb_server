package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
)

// UserInfoView 用户信息
// @Tags 用户管理
// @Summary 用户信息
// @Description 用户信息
// @Router /api/user_info [get]
// @Param token header string  true  "token"
// @Produce json
// @Success 200 {object} res.Response{data=models.UserModel}
func (UserApi) UserInfoView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	// 查询用户是否存在
	var userModel models.UserModel
	count := global.DB.Take(&userModel, claims.UserID).RowsAffected
	if count == 0 {
		res.FailWithMsg("用户不存在", c)
		return
	}
	res.OKWithData(userModel, c)
}
