package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/service/common"
	"gvb_server/utils/res"
	"strings"
)

// AdvertListView 查询广告列表
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	isShow := true
	referer := c.GetHeader("GVB_REFERER")
	if strings.Contains(referer, "admin") {
		isShow = false
	}
	list, count, err := common.CommList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		res.FailWithMsg("查询广告列表失败", c)
		return
	}
	res.OKWithList(list, count, c)
}
