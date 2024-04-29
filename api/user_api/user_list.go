package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/service/common"
	"gvb_server/utils/desens"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
)

type UserResponse struct {
	models.UserModel
	RoleID int `json:"role_id"`
}

type UserListRequest struct {
	models.PageInfo
	Role int `json:"role"`
}

// UserListView 查询用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Param data query models.PageInfo  false  "查询参数"
// @Param token header string  true  "token"
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.UserModel]}
func (UserApi) UserListView(c *gin.Context) {
	var cr UserListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		global.Log.Error(err)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var userResponseList []UserResponse
	list, count, err := common.CommList(models.UserModel{Role: ctype.Role(cr.Role)}, common.Option{
		PageInfo: cr.PageInfo,
		Debug:    true,
	})
	if err != nil {
		res.FailWithMsg("查询用户列表失败", c)
		return
	}
	for _, model := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			model.UserName = ""
		}
		// 手机号脱敏
		model.Tel = desens.DesensitizationPhone(model.Tel)
		// 邮箱脱敏
		model.Email = desens.DesensitizationEmail(model.Email)
		userResponseList = append(userResponseList, UserResponse{
			UserModel: model,
			RoleID:    int(model.Role),
		})
	}
	res.OKWithList(userResponseList, count, c)
}
