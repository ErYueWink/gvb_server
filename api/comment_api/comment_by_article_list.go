package comment_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type CommentByArticleListRequest struct {
	models.PageInfo
	Title string `json:"title" form:"title"`
}

type CommentByArticleListResponse struct {
	Title string `json:"title"`
	ID    string `json:"id"`
	Count int64  `json:"count"`
}

// CommentByArticleListView 有评论的文章列表
// @Tags 评论管理
// @Summary 有评论的文章列表
// @Description 有评论的文章列表
// @Param id path string  true  "id"
// @Param data query CommentByArticleListRequest  true  "参数"
// @Router /api/comments/articles [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[CommentByArticleListResponse]}
func (CommentApi) CommentByArticleListView(c *gin.Context) {
	var cr CommentByArticleListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var count int64
	global.DB.Model(models.CommentModel{}).Group("article_id").Count(&count)
	offset := (cr.Page - 1) * cr.Limit
	type T struct {
		ArticleID string `json:"article_id"`
		Count     int64  `json:"count"`
	}
	var _list []T
	global.DB.Model(models.CommentModel{}).
		Group("article_id").Order("count desc").Select("article_id", "count(id) as count").
		Offset(offset).Limit(cr.Limit).Scan(&_list)
	var articleIDMap map[string]int64
	var articleIDList []interface{}
	for _, t := range _list {
		articleIDMap[t.ArticleID] = t.Count
		articleIDList = append(articleIDList, t.ArticleID)
	}
	// 根据文章ID查询判断文章是否真实存在
	result, err := global.EsClient.Search(models.ArticleModel{}.Index()).
		Query(elastic.NewTermsQuery("_id", articleIDList...)).
		Size(10000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("ES查询失败", c)
		return
	}
	// 响应数据
	var commentByArticleListResponse = make([]CommentByArticleListResponse, 0)
	for _, hit := range result.Hits.Hits {
		var articleModel models.ArticleModel
		err = json.Unmarshal(hit.Source, &articleModel)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		articleModel.ID = hit.Id
		commentByArticleListResponse = append(commentByArticleListResponse, CommentByArticleListResponse{
			Title: articleModel.Title,
			ID:    hit.Id,
			Count: articleIDMap[hit.Id],
		})
	}
	res.OKWithList(commentByArticleListResponse, count, c)
	return
}
