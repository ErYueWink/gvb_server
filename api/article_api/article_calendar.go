package article_api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// ArticleCalendarView 以天为单位进行聚合查询
func (ArticleApi) ArticleCalendarView(c *gin.Context) {
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")
	result, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewBoolQuery()).
		Size(0).
		Aggregation("calendar", agg).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg("查询日历失败", c)
		return
	}
	res.OKWithData(string(result.Aggregations["calendar"]), c)
}
