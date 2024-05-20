package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/res"
)

// ArticleContentByIDView 获取文章正文
// @Tags 文章管理
// @Summary 获取文章正文
// @Description 获取文章正文
// @Param id path int  true  "id"
// @Router /api/articles/content/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleContentByIDView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	result, err := global.EsClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(cr.ID).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询文章正文失败", c)
		return
	}
	var model models.ArticleModel
	err = json.Unmarshal(result.Source, &model)
	if err != nil {
		global.Log.Error(err)
		return
	}
	// 增加文章浏览量
	redis_ser.NewArticleLook().Set(cr.ID)
	// 返回正文内容
	res.OKWithData(model.Content, c)
}
