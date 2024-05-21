package es_ser

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
)

type SearchData struct {
	Key   string `json:"key"`   // 文章关联ID
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
}

// DeleteFullTextByArticleID 删除文章数据
func DeleteFullTextByArticleID(id string) {
	termQuery := elastic.NewTermQuery("key", id)
	result, err := global.EsClient.
		DeleteByQuery().
		Index(models.FullTextModel{}.Index()).
		Query(termQuery).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(fmt.Sprintf("共删除%d个数据", result.Deleted))
}
