package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// AdvertRemoveView 批量删除广告
// @Tags 广告管理API
// @Summary 批量删除广告
// @Description 批量删除广告
// @Param token header string  true  "token"
// @Param data body models.RemoveRequest    true  "广告id列表"
// @Router /api/advert [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("广告不存在", c)
		return
	}
	err = global.DB.Delete(&advertList).Error
	if err != nil {
		res.FailWithMsg("删除广告失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("共删除%d个广告", count), c)
}
