package es_ser

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"strings"
)

type Option struct {
	models.PageInfo
	Field []string
	Tag   string // java go python
}

// GetForm
func (option *Option) GetForm() int {
	if option.Limit == 0 {
		option.Limit = 10
	}
	form := (option.Page - 1) * option.Limit
	if form < 0 {
		form = 0
	}
	return form
}

// CommonList ES通用列表询
func CommonList(option Option) (list []models.ArticleModel, count int64, err error) {
	boolSearch := elastic.NewBoolQuery()
	if option.Key != "" {
		boolSearch.Must(elastic.NewMultiMatchQuery(option.Key, option.Field...))
	}
	if option.Tag != "" {
		boolSearch.Must(elastic.NewMultiMatchQuery(option.Tag, "tags"))
	}
	type SortField struct {
		Field     string
		Ascending bool // 正序 asc
	}
	sortField := SortField{
		Field:     "created_at",
		Ascending: false,
	}
	if option.Sort != "" { // created_at asc
		_list := strings.Split(option.Sort, " ")
		if len(_list) == 2 && (_list[1] == "asc" || _list[1] == "desc") {
			sortField.Field = _list[0]
			if _list[1] == "asc" {
				sortField.Ascending = true
			}
			if _list[1] == "desc" {
				sortField.Ascending = false
			}
		}
	}
	res, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		From(option.GetForm()).
		Size(option.Limit).
		Sort(sortField.Field, sortField.Ascending).
		Do(context.Background())
	count = res.Hits.TotalHits.Value
	var articleList []models.ArticleModel
	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err)
			continue
		}
		err = json.Unmarshal(data, &article)
		if err != nil {
			logrus.Error(err)
			continue
		}
		article.ID = hit.Id
		articleList = append(articleList, article)
	}
	return articleList, count, err
}

// CommonDetail 通用详情查询
func CommonDetail(id string) (model models.ArticleModel, err error) {
	res, err := global.EsClient.
		Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	data, err := res.Source.MarshalJSON()
	if err != nil {
		logrus.Error(err)
		return
	}
	err = json.Unmarshal(data, &model)
	if err != nil {
		logrus.Error(err)
		return
	}
	model.ID = res.Id
	return model, err
}

// CommonDetailByTitle 根据标题查询文章详情
func CommonDetailByTitle(key string) (model models.ArticleModel, err error) {
	res, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", key)).Size(1).
		Size(1).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return model, errors.New("查询文章详情失败")
	}
	hit := res.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &model)
	if err != nil {
		logrus.Error(err)
		return model, errors.New("查询文章详情失败")
	}
	model.ID = hit.Id
	return model, err
}
