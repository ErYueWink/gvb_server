package comment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/res"
)

type CommentIDRequest struct {
	ID uint `json:"id" uri:"id"`
}

// CommentDigg 评论点赞
// @Tags 评论管理
// @Summary 评论点赞
// @Description 评论点赞
// @Param id path int  true  "id"
// @Router /api/comments/digg/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (CommentApi) CommentDigg(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	// 查询评论是否存在
	var model models.CommentModel
	count := global.DB.Take(&model, cr.ID).RowsAffected
	if count == 0 {
		res.FailWithMsg("评论不存在", c)
		return
	}
	// 评论存在 添加评论点赞数量
	err = redis_ser.NewCommentDigg().Set(fmt.Sprintf("%d", cr.ID))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("评论点赞失败", c)
		return
	}
	// 将评论点赞数量同步到数据库中
	err = global.DB.Model(&model).
		Update("digg_count", redis_ser.NewCommentDigg().Get(fmt.Sprintf("%d", cr.ID))).
		Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("同步点赞数量失败", c)
		return
	}
	res.OKWithMsg("评论点赞成功", c)
}
