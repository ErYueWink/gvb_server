package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/es_ser"
	"gvb_server/utils/res"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag string `form:"tag" json:"tag"`
}

// ArticleListView 文章列表
// @Tags 文章管理
// @Summary 文章列表
// @Description 文章列表
// @Param data query ArticleSearchRequest   false  "表示多个参数"
// @Param token header string  false  "token"
// @Router /api/articles [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.ArticleModel]}
func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	list, count, err := es_ser.ESComList(es_ser.Option{
		PageInfo: cr.PageInfo,
		Tag:      cr.Tag,
		Fields:   []string{"title", "abstract", "content"},
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询文章列表失败", c)
		return
	}
	res.OKWithList(list, count, c)
}
