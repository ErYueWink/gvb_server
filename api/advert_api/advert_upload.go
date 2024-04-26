package advert_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// AdvertUpdateView 修改广告
// @Tags 广告管理API
// @Summary 更新广告
// @Param token header string  true  "token"
// @Description 更新广告
// @Param data body AdvertRequest    true  "广告的一些参数"
// @Param id path int true "id"
// @Router /api/advert/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	// 参数绑定
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var advertModel models.AdvertModel
	count := global.DB.Take(&advertModel, id).RowsAffected
	if count == 0 {
		res.FailWithMsg("广告不存在", c)
		return
	}
	// 结构体转Map
	advertMap := structs.Map(&cr)
	// 修改广告
	err = global.DB.Model(&advertModel).Updates(advertMap).Error
	if err != nil {
		res.FailWithMsg("修改广告失败", c)
		return
	}
	res.OKWithMsg("修改广告成功", c)

}
