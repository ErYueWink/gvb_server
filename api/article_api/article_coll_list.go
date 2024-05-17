package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/common"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
)

type CollResponse struct {
	models.ArticleModel
	CreatedAt string `json:"created_at"`
}

// ArticleCollListView 用户收藏的文章列表
// @Tags 文章管理
// @Summary 用户收藏的文章列表
// @Description 用户收藏的文章列表
// @Param data query models.PageInfo  true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles/collects [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[CollResponse]}
func (ArticleApi) ArticleCollListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	// 查询用户收藏的文章列表
	var articleIDList []interface{}
	list, count, err := common.CommList(models.UserCollectModel{UserID: claims.UserID}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询用户收藏列表失败", c)
		return
	}
	var collMap map[string]string
	// 记录文章收藏时间
	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		collMap[model.ArticleID] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}
	// 构造查询条件
	query := elastic.NewTermsQuery("_id", articleIDList...)
	// 查询文章
	result, err := global.EsClient.Search(models.ArticleModel{}.Index()).
		Query(query).Size(1000).Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询文章列表失败", c)
		return
	}
	var collList = make([]CollResponse, 0)
	for _, hit := range result.Hits.Hits {
		var articleColl models.ArticleModel
		err = json.Unmarshal(hit.Source, &articleColl)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		articleColl.ID = hit.Id
		articleColl.Content = ""
		collList = append(collList, CollResponse{
			articleColl,
			collMap[hit.Id],
		})
	}
	res.OKWithList(collList, count, c)
}
