package article_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/service/es_ser"
	"gvb_server/utils/res"
)

type ESIDRequest struct {
	ID string `json:"id" uri:"id" form:"id"`
}

// ArticleDetailView 文章详情
// @Tags 文章管理
// @Summary 文章详情
// @Description 文章详情
// @Param id path string  true  "id"
// @Router /api/articles/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=models.ArticleModel}
func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	articleModel, err := es_ser.CommonDetail(cr.ID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(fmt.Sprintf("ID为%s的文章查询失败", cr.ID), c)
		return
	}
	res.OKWithData(articleModel, c)
}
