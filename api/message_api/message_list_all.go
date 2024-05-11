package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/common"
	"gvb_server/utils/res"
)

// MessageListAllView 消息列表
// @Tags 消息管理
// @Summary 消息列表
// @Description 消息列表
// @Router /api/messages_all [get]
// @Param token header string  true  "token"
// @Param data query models.PageInfo    false  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.MessageModel]}
func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBind(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.CommList(models.MessageModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询消息列表失败", c)
		return
	}
	res.OKWithList(list, count, c)

}
