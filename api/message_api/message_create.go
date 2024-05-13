package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
)

type MessageRequest struct {
	RevUserID uint   `json:"rev_user_id" binding:"required" msg:"接收者ID"`
	Content   string `json:"content" binding:"required" msg:"消息内容"`
}

// MessageCreateView 发布消息
// @Tags 消息管理
// @Summary 发布消息
// @Description 发布消息
// @Param data body MessageRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/messages [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (MessageApi) MessageCreateView(c *gin.Context) {
	var cr MessageRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var sendUser, revUser models.UserModel
	// 发送者ID就是当前登录的用户ID
	// 获取cookie中的载荷信息
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	err = global.DB.Take(&sendUser, claims.UserID).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("发送者不存在", c)
		return
	}
	err = global.DB.Take(&revUser, cr.RevUserID).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("接收者不存在", c)
		return
	}
	// 保存消息记录
	err = global.DB.Create(&models.MessageModel{
		SendUserID:       sendUser.ID,
		SendUserNickName: sendUser.NickName,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        revUser.ID,
		RevUserNickName:  revUser.NickName,
		RevUserAvatar:    revUser.Avatar,
		Content:          cr.Content,
		IsRead:           false,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("保存消息记录失败", c)
		return
	}
	res.OKWithMsg("发送消息成功", c)
}
