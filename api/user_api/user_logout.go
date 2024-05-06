package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/service/user_ser"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
)

// LogoutView 用户注销
// @Tags 用户管理
// @Summary 用户注销
// @Description 用户注销
// @Param token header string  true  "token"
// @Router /api/logout [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) LogoutView(c *gin.Context) {
	// 获取载荷信息
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	token := c.Request.Header.Get("token")
	// 用户注销
	err := user_ser.UserService{}.Logout(claims, token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("用户注销失败", c)
		return
	}
	res.OKWithMsg("用户注销成功", c)

}
