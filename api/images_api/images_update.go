package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type ImagesUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请输入图片ID"`
	Name string `json:"name" binding:"required" msg:"请输入图片名称"`
}

// ImagesUpdateView 编辑图片
func (ImagesApi) ImagesUpdateView(c *gin.Context) {
	var cr ImagesUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, cr.ID).Error
	if err != nil {
		res.FailWithMsg("图片不存在", c)
		return
	}
	// 图片存在修改图片
	err = global.DB.Model(&bannerModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMsg("图片修改失败", c)
		return
	}
	res.OKWithMsg("图片修改成功", c)
}
