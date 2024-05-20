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

type ArticleCategoryResponse struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

// ArticleCategoryListView 文章分类列表
// @Tags 文章管理
// @Summary 文章分类列表
// @Description 文章分类列表
// @Router /api/categories [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]CategoryResponse}
func (ArticleApi) ArticleCategoryListView(c *gin.Context) {
	type T struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}
	agg := elastic.NewTermsAggregation().Field("category")
	result, err := global.EsClient.Search(models.ArticleModel{}.Index()).
		Query(elastic.NewBoolQuery()).
		Aggregation("categorys", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg("查询文章分类列表失败", c)
		return
	}
	byteData := result.Aggregations["categorys"]
	var data T
	_ = json.Unmarshal(byteData, &data)
	var categoryResponse = make([]ArticleCategoryResponse, 0)
	for _, bucket := range data.Buckets {
		categoryResponse = append(categoryResponse, ArticleCategoryResponse{
			Label: bucket.Key,
			Value: bucket.DocCount,
		})
	}
	res.OKWithData(categoryResponse, c)
}
