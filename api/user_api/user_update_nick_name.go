package user_api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
	"strings"
)

type UserUpdateNickNameRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Sign     string `json:"sign" structs:"sign"`
	Link     string `json:"link" structs:"link"`
	Avatar   string `json:"avatar" structs:"avatar"`
}

// UserUpdateNickNameView 修改当前登录人的昵称，签名，链接
// @Tags 用户管理
// @Summary 修改当前登录人的昵称，签名，链接
// @Description 修改当前登录人的昵称，签名，链接
// @Router /api/user_info [put]
// @Param token header string  true  "token"
// @Param data body UserUpdateNicknameRequest  true  "昵称，签名，链接"
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdateNickNameView(c *gin.Context) {
	var cr UserUpdateNickNameRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("参数绑定失败", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.CustomClaims)
	var newMaps = map[string]interface{}{}
	maps := structs.Map(cr)
	// 只修改有值的
	for k, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[k] = val
		}
	}
	// 根据用户ID查询判断要修改的用户是否存在
	var userModel models.UserModel
	count := global.DB.Take(&userModel, claims.UserID).RowsAffected
	if count == 0 {
		res.FailWithMsg(fmt.Sprintf("用户ID为%d的用户不存在", claims.UserID), c)
		return
	}
	// 如果修改的是头像则判断用户的注册来源
	_, ok := newMaps["avatar"]
	if ok && userModel.SignStatus != ctype.SignEmail { // 只有注册来源是邮箱的才可以修改头像
		delete(newMaps, "avatar")
	}
	err = global.DB.Model(&userModel).Updates(newMaps).Error
	if err != nil {
		res.FailWithMsg(fmt.Sprintf("修改%s的个人信息失败", claims.NickName), c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("修改%s的个人信息成功", claims.NickName), c)
}
