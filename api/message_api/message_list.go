package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
	"time"
)

type Message struct {
	SendUserID       uint      `json:"send_user_id"` // 发送人id
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `json:"rev_user_id"` // 接收人id
	RevUserNickName  string    `json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	Content          string    `json:"content"`       // 消息内容
	CreatedAt        time.Time `json:"created_at"`    // 最新的消息时间
	MessageCount     int       `json:"message_count"` // 消息条数
}

type MessageGroup map[uint]*Message

// MessageListView 用户与其他人的消息列表
// @Tags 消息管理
// @Summary 用户与其他人的消息列表
// @Description 用户与其他人的消息列表
// @Router /api/messages [get]
// @Param token header string  true  "token"
// @Produce json
// @Success 200 {object} res.Response{data=[]Message}
func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var messageGroup = MessageGroup{}
	var messageList []models.MessageModel
	var messages = make([]Message, 0)
	err := global.DB.Order("created_at asc").
		Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询消息列表失败", c)
		return
	}
	for _, message := range messageList {
		messageInfo := Message{
			SendUserID:       message.SendUserID,
			SendUserNickName: message.SendUserNickName,
			SendUserAvatar:   message.SendUserAvatar,
			RevUserID:        message.RevUserID,
			RevUserNickName:  message.RevUserNickName,
			RevUserAvatar:    message.RevUserAvatar,
			Content:          message.Content,
			CreatedAt:        message.CreatedAt,
			MessageCount:     1,
		}
		// 判断是否为一组
		idNum := message.SendUserID + message.RevUserID
		val, ok := messageGroup[idNum]
		if !ok {
			messageGroup[idNum] = &messageInfo
			continue
		}
		messageInfo.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &messageInfo
	}
	for _, message := range messageGroup {
		messages = append(messages, *message)
	}
	res.OKWithData(messages, c)
	return
}
