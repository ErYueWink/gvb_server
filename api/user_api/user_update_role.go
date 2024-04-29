package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type UserUpdateRoleRequest struct {
	UserID   int    `json:"user_id" binding:"required" msg:"请输入"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数非法"`
}

// UserUpdateRoleView 用户权限变更
// @Tags 用户管理
// @Summary 用户权限变更
// @Description 用户权限变更
// @Param token header string  true  "token"
// @Param data body UserRole  true  "查询参数"
// @Router /api/user_role [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserUpdateRoleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil { // 参数绑定失败
		res.FailWithError(err, &cr, c)
		return
	}
	// 根据用户ID查询
	var userModel models.UserModel
	count := global.DB.Take(&userModel, "id=?", cr.UserID).RowsAffected
	if count == 0 {
		res.FailWithMsg("用户不存在", c)
		return
	}
	// 修改用户
	err = global.DB.Model(&userModel).Updates(map[string]interface{}{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		res.FailWithMsg("修改用户失败", c)
		return
	}
	res.OKWithMsg("修改用户成功", c)
}
