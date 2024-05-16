package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/es_ser"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
)

// ArticleCollCreateView 用户收藏文章，或取消收藏
// @Tags 文章管理
// @Summary 用户收藏文章，或取消收藏
// @Description 用户收藏文章，或取消收藏
// @Param data body models.ESIDRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles/collects [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleCollCreateView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	articleModel, err := es_ser.CommonDetail(cr.ID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("文章不存在", c)
		return
	}
	// 获取当前登录用户ID
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	// 判断用户有没有收藏该文章
	var coll models.UserCollectModel
	count := global.DB.Find(&coll, "user_id = ? and article_id = ?", claims.UserID, cr.ID).RowsAffected
	num := -1
	// 用户没有收藏文章
	if count == 0 {
		err = global.DB.Create(models.UserCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		}).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("用户收藏文章失败", c)
			return
		}
		// 文章收藏数+1
		num = 1
	}
	// 取消收藏 如果查询到这篇文章则进行删除
	global.DB.Delete(&coll)
	// 更新文章收藏总数
	err = es_ser.ArticleUpdate(cr.ID, map[string]any{
		"collects_count": articleModel.CollectsCount + num,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("收藏文章/取消收藏成功", c)
		return
	}
	if num == 1 {
		res.OKWithMsg("收藏文章成功", c)
		return
	} else {
		res.OKWithMsg("取消收藏成功", c)
		return
	}

}
