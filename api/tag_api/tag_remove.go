package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// TagRemoveView 标签删除
// @Tags 标签管理
// @Summary 标签删除
// @Description 标签删除
// @Param data body models.RemoveRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/tags [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		global.Log.Error(err)
		res.FailWithMsg("标签不存在", c)
		return
	}
	// 标签存在的情况下删除标签
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO 该标签被文章关联怎么办
		err = global.DB.Delete(&tagList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		res.FailWithMsg("删除标签失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("共删除%d个标签", count), c)
}
