package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// FullTextSearchView 全文搜索列表
// @Tags 文章管理
// @Summary 全文搜索列表
// @Description 全文搜索列表
// @Param data query models.PageInfo   false  "表示多个参数"
// @Router /api/articles/text [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.FullTextModel]}
func (ArticleApi) FullTextSearchView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	boolQuery := elastic.NewBoolQuery()
	if cr.Key == "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}
	if cr.Limit == 0 {
		cr.Limit = 10
	}
	if cr.Page == 0 {
		cr.Page = 1
	}
	form := (cr.Page - 1) * cr.Limit
	// 搜索
	result, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).Query(boolQuery).
		From(form).Size(cr.Limit).
		Highlight(elastic.NewHighlight().Field("body")).
		Do(context.Background())
	if err != nil {
		return
	}
	// 搜索结果总条数
	count := result.Hits.TotalHits.Value
	var fullTextList = make([]models.FullTextModel, 0)
	for _, hit := range result.Hits.Hits {
		var fullText models.FullTextModel
		err = json.Unmarshal(hit.Source, &fullText)
		if err != nil {
			logrus.Error(err)
			continue
		}
		val, ok := hit.Highlight["body"]
		if ok {
			fullText.Body = val[0]
		}
		fullTextList = append(fullTextList, fullText)
	}
	res.OKWithList(fullTextList, count, c)

}
