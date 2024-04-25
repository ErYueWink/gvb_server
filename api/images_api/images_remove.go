package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// ImagesRemoveApi 删除图片
func (ImagesApi) ImagesRemoveApi(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var bannerList []models.BannerModel
	count := global.DB.Select("id").Find(&bannerList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("没有该图片", c)
		return
	}
	// 删除图片
	err = global.DB.Delete(&bannerList).Error
	if err != nil {
		res.FailWithMsg("删除图片失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("共删除%d张图片", count), c)
}
