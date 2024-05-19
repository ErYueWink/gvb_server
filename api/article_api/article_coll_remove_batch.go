package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/es_ser"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
)

// ArticleCollBatchRemoveView 用户取消收藏文章
// @Tags 文章管理
// @Summary 用户取消收藏文章
// @Description 用户取消收藏文章
// @Param data body models.ESIDListRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles/collects [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleCollBatchRemoveView(c *gin.Context) {
	var cr models.ESListIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var articleIDList []string
	var collList []models.ArticleModel
	// 查询用户要取消收藏的文章
	global.DB.Find(&collList, "user_id = ? or article_id in ?", claims.UserID, cr.IDList).
		Select("article_id").Scan(&articleIDList)
	if len(articleIDList) == 0 {
		res.FailWithMsg("请求非法", c)
		return
	}
	var IDList []interface{}
	for _, articleID := range articleIDList {
		IDList = append(IDList, articleID)
	}
	// 定义查询条件
	query := elastic.NewTermsQuery("_id", IDList...)
	// 查询ES
	result, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Size(1000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("取消收藏文章成功", c)
		return
	}
	for _, hit := range result.Hits.Hits {
		var articleModel models.ArticleModel
		err = json.Unmarshal(hit.Source, &articleModel)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		// 文章的收藏数减1
		count := articleModel.CollectsCount - 1
		// 修改文章收藏数
		err = es_ser.ArticleUpdate(hit.Id, map[string]any{
			"collects_count": count,
		})
		if err != nil {
			global.Log.Error(err)
			continue
		}
	}
	// 删除中间表数据
	global.DB.Delete(&collList)
	res.OKWithMsg(fmt.Sprintf("共取消收藏%d篇文章", len(articleIDList)), c)
}
