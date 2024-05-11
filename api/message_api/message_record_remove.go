package message_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// MessageRecordRemoveView 删除用户的消息记录
// @Tags 消息管理
// @Summary 删除用户的消息记录
// @Description 删除用户的消息记录
// @Router /api/message_users [delete]
// @Param token header string  true  "token"
// @Param data body models.RemoveRequest   true  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{]}
func (MessageApi) MessageRecordRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var messageList []models.MessageModel
	err = global.DB.Find(&messageList, cr.IDList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询用户消息记录失败", c)
		return
	}
	if len(messageList) > 0 {
		err = global.DB.Delete(&messageList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("删除用户消息记录失败", c)
			return
		}
	}
	res.OKWithMsg(fmt.Sprintf("共删除%d条记录", len(messageList)), c)

}
