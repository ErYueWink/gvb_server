package comment_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type CommentListRequest struct {
	ArticleID string `form:"article_id"`
}

// CommentListView 文章下的评论列表
// @Tags 评论管理
// @Summary 文章下的评论列表
// @Description 文章下的评论列表
// @Param id path string  true  "id"
// @Router /api/comments/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]models.CommentModel}
func (CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	rootCommentList := FindArticleCommentList(cr.ArticleID)
	res.OKWithData(filter.Select("c", rootCommentList), c)
	return
}

// FindArticleCommentList 查询文章下的所有评论
func FindArticleCommentList(article string) (RootComment []*models.CommentModel) {
	/**
	  1. 查询所有根评论
	  2. 遍历每个根评论 查询出根评论下的所有子评论
	*/
	global.DB.Find(&RootComment, "article_id = ? and parent_comment_id is not null", article)
	for _, model := range RootComment {
		var subCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		model.SubComments = subCommentList
	}
	return
}

// FindSubComment 递归查询评论下所有的子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	// 预加载子评论以及用户数据
	global.DB.Preload("SubComments.User").Find(&model)
	for _, comment := range model.SubComments {
		*subCommentList = append(*subCommentList, comment)
		// 递归 查询评论下的所有子评论
		FindSubComment(comment, subCommentList)
	}
	return
}

// FindSubCommentCount 递归查评论下的子评论数
func FindSubCommentCount(model models.CommentModel) []models.CommentModel {
	var subCommentList []models.CommentModel
	FindSubCommentList(model, &subCommentList)
	return subCommentList
}

// FindSubCommentList 递归查询评论下的子评论列表
func FindSubCommentList(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments").Find(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(model, subCommentList)
	}
	return
}
