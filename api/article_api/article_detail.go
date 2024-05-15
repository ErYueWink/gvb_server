package article_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/service/es_ser"
	"gvb_server/utils/res"
)

type ESIDRequest struct {
	ID string `json:"id" uri:"id" form:"id"`
}

// ArticleDetailView 根据ID查询文章详情文章详情
// @Tags 文章管理
// @Summary 文章详情
// @Description 文章详情
// @Param id path string  true  "id"
// @Router /api/articles/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=models.ArticleModel}
func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	articleModel, err := es_ser.CommonDetail(cr.ID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(fmt.Sprintf("ID为%s的文章查询失败", cr.ID), c)
		return
	}
	res.OKWithData(articleModel, c)
}

type ESTitleRequest struct {
	Title string `json:"title" uri:"title" form:"title"`
}

// ArticleDetailByTitleView 根据ID查询文章详情文章详情
// @Tags 文章管理2
// @Summary 文章详情2
// @Description 文章详情2
// @Param title path string  true  "title"
// @Router /api/articles/detail [get]
// @Produce json
// @Success 200 {object} res.Response{data=models.ArticleModel}
func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ESTitleRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	articleModel, err := es_ser.CommonDetailByTitle(cr.Title)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询文章详情失败", c)
	}
	res.OKWithData(articleModel, c)

}
