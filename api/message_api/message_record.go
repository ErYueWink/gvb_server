package message_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required"`
}

// MessageRecordView 用户的消息记录
// @Tags 消息管理
// @Summary 用户的消息记录
// @Description 用户的消息记录
// @Router /api/messages_record [post]
// @Param token header string  true  "token"
// @Param data body MessageRecordRequest  true  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{data=[]models.MessageModel}
func (MessageApi) MessageRecordView(c *gin.Context) {
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	// 查询用户所有消息
	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)
	err = global.DB.Find(&_messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询用户消息列表失败", c)
		return
	}
	for _, model := range _messageList {
		if model.SendUserID == cr.UserID || model.RevUserID == cr.UserID {
			messageList = append(messageList, model)
		}
	}
	// 用户查看和谁的消息记录，那条消息记录就变为已读
	for _, model := range messageList {
		err = global.DB.Model(&model).Update("is_read", true).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Info(fmt.Sprintf("用户已读和%d用户的消息", cr.UserID))
	}
	res.OKWithData(messageList, c)
}
