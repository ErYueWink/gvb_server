package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入广告标题"  struct:"title""` // 广告标题
	Href   string `json:"href" binding:"required" msg:"请输入广告跳转链接" struct:"href"`   // 广告跳转链接
	Images string `json:"images" binding:"required" msg:"请输入广告图片" struct:"images"` // 广告图片
	IsShow bool   `json:"is_show" struct:"is_show"`                                // 是否显示
}

// AdvertCreateView 创建广告
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var advert models.AdvertModel
	count := global.DB.Take(&advert, "title = ?", cr.Title).RowsAffected
	if count != 0 {
		res.FailWithMsg(fmt.Sprintf("广告%s重复", cr.Title), c)
		return
	}
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		res.FailWithMsg("创建广告失败", c)
		return
	}
	res.OKWithMsg("创建广告成功", c)
}
