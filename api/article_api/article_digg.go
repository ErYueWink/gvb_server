package article_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/es_ser"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/res"
)

// ArticleDiggView 文章点赞
// @Tags 文章管理
// @Summary 文章点赞
// @Description 文章点赞
// @Param data body models.ESIDRequest   true  "表示多个参数"
// @Router /api/articles/digg [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleDiggView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	// 查询ES判断文章是否真实存在
	result, err := global.EsClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(cr.ID).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(fmt.Sprintf("文章%s不存在", cr.ID), c)
		return
	}
	// 点赞文章
	redis_ser.NewArticleDigg().Set(result.Id)
	// 同步点赞数量到ES
	err = es_ser.ArticleUpdate(result.Id, map[string]any{
		"digg_count": redis_ser.NewArticleDigg().Get(result.Id),
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("同步点赞数量失败", c)
		return
	}
	res.OKWithMsg("文章点赞成功", c)
}
