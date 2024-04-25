package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type ImagesNameResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// ImagesNameListView 查询图片名称列表
func (ImagesApi) ImagesNameListView(c *gin.Context) {
	var response []ImagesNameResponse
	err := global.DB.Model(&models.BannerModel{}).Select("id", "name", "path").Scan(&response).Error
	if err != nil {
		res.FailWithMsg("查询图片名称列表失败", c)
		return
	}
	res.OKWithData(response, c)
}
