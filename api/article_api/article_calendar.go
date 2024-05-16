package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
	"time"
)

type BucketsType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int    `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

type CalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

// 2024-05-16 5
var DateCount = map[string]int{}

// ArticleCalendarView 文章日历
func (ArticleApi) ArticleCalendarView(c *gin.Context) {
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")
	now := time.Now()
	aYear := now.AddDate(-1, 0, 0)
	format := "2006-01-02 15:04:05"
	query := elastic.NewRangeQuery("created_at").Gte(aYear.Format(format)).Lte(now.Format(format))
	result, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).Query(query).
		Size(0).Aggregation("calendar", agg).Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询文章日历失败", c)
		return
	}
	var buckets BucketsType
	_ = json.Unmarshal(result.Aggregations["calendar"], &buckets)
	for _, bucket := range buckets.Buckets {
		Time, _ := time.Parse(format, bucket.KeyAsString)
		DateCount[Time.Format("2006-01-02")] = bucket.DocCount
	}
	days := int(now.Sub(aYear).Hours() / 24)
	var calendarResponse []CalendarResponse
	for i := 0; i < days; i++ {
		day := aYear.AddDate(0, 0, i).Format("2006-01-02")
		count, _ := DateCount[day]
		calendarResponse = append(calendarResponse, CalendarResponse{
			Date:  day,
			Count: count,
		})
	}
	res.OKWithData(calendarResponse, c)
}
