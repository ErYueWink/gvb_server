package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/service/user_ser"
	"gvb_server/utils/res"
)

type UserCreateRequest struct {
	Username string     `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string     `json:"password" binding:"password" msg:"请输入密码"`
	NickName string     `json:"nick_name" binding:"nickname" msg:"请输入用户昵称"`
	Role     ctype.Role `json:"role" binding:"role" msg:"请输入用户权限"`
}

// UserCreateView 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户
// @Param data body UserCreateRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/users [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserCreateView(c *gin.Context) {
	var cr models.UserModel
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	err = user_ser.UserService{}.CreateUser(cr.UserName, cr.Password, cr.NickName, "", cr.Role, c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("创建用户失败", c)
		return
	}
	res.OKWithMsg("创建用户成功", c)
}
