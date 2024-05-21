package comment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/es_ser"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择要评论的文章"`
	Content         string `json:"content" binding:"required" msg:"请输入文章内容"`
	ParentCommentID *uint  `json:"parent_comment_id"`
}

// CommentCreateRequest 发布评论
func (CommentApi) CommentCreateRequest(c *gin.Context) {
	var cr CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	// 查询文章详情判断文章是否存在
	_, err = es_ser.CommonDetail(cr.ArticleID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(fmt.Sprintf("编号为%s的文章不存在", cr.ArticleID), c)
		return
	}
	var parentComment models.CommentModel
	// 判断前端有没有传父评论ID
	if cr.ParentCommentID != nil {
		count := global.DB.Take(&parentComment, cr.ParentCommentID).RowsAffected
		if count == 0 {
			res.FailWithMsg("父评论不存在", c)
			return
		}
		// 判断父评论的文章和当前文章是否一致
		if parentComment.ArticleID != cr.ArticleID {
			res.FailWithMsg("评论文章不一致", c)
			return
		}
	}
	// 发布评论
	err = global.DB.Create(&models.CommentModel{
		ArticleID:       cr.ArticleID,
		Content:         cr.Content,
		ParentCommentID: cr.ParentCommentID,
	}).Error
	if err != nil {
		res.FailWithMsg("发布评论失败", c)
		return
	}
	// 发布评论之后使父评论下的子评论数+1
	global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + 1"))
	// 文章评论数同步到缓存中
	err = redis_ser.NewArticleComment().Set(cr.ArticleID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("评论数量同步到Redis失败", c)
		return
	}
	// 同步文章评论数到ES
	err = es_ser.ArticleUpdate(cr.ArticleID, map[string]any{
		"comment_count": redis_ser.NewArticleComment().Get(cr.ArticleID),
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("评论数量同步到ES失败", c)
		return
	}
	res.OKWithMsg("发布评论成功", c)
	return
}
