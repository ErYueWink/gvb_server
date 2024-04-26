package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/service/common"
	"gvb_server/utils/res"
	"strings"
)

// AdvertListView 查询广告列表
// @Tags 广告管理API
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/advert [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
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
