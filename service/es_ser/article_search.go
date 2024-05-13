package es_ser

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"strings"
)

type Option struct {
	models.PageInfo
	Fields []string
	Tag    string
}

func (o *Option) GetForm() int {
	if o.Limit == 0 {
		o.Limit = 10
	}
	form := (o.Page - 1) * o.Limit
	if form < 0 {
		form = 0
	}
	return form
}

// ESComList ES通用查询
func ESComList(option Option) (list []models.ArticleModel, count int64, err error) {
	// 构造复杂查询条件
	boolSearch := elastic.NewBoolQuery()
	// 如果关键字不为空，则根据关键字进行范围查询
	if option.Key != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.Key, option.Fields...),
		)
	}
	// 如果标签不为空，则根据标签进行查询
	if option.Tag != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.Tag, "tags"))
	}
	// 排序
	type SortField struct {
		Field     string
		Ascending bool // 默认asc
	}
	// 设置默认值
	sortField := SortField{
		Field:     "created_at",
		Ascending: false,
	}
	if option.Sort != "" {
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
	// 搜索
	res, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		From(option.GetForm()).Size(option.Limit).
		Sort(sortField.Field, sortField.Ascending).
		Do(context.Background())
	// 获取总条数
	count = res.Hits.TotalHits.Value
	var articleList []models.ArticleModel
	// 遍历查询结果集
	for _, hit := range res.Hits.Hits {
		var articleModel models.ArticleModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err)
			continue
		}
		err = json.Unmarshal(data, &articleModel)
		if err != nil {
			logrus.Error(err)
			continue
		}
		articleModel.ID = hit.Id
		articleList = append(articleList, articleModel)
	}
	return articleList, count, err
}
