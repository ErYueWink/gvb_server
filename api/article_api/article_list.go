package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/es_ser"
	"gvb_server/utils/res"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag string `json:"tag" form:"tag"`
}

// ArticleListView 文章搜索接口
func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	list, count, err := es_ser.CommonList(es_ser.Option{
		PageInfo: cr.PageInfo,
		Field:    []string{"title", "abstract", "content"},
		Tag:      cr.Tag,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询文章列表失败", c)
		return
	}
	res.OKWithList(filter.Omit("list", list), count, c)
}
