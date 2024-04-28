package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"` // 切换的时间，单位秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`     // 切换的时间，单位秒
	Sort          int         `json:"sort" structs:"sort"`                   // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`           // 具体图片的顺序
}

// MenuCreateView 发布菜单
// @Tags 菜单管理
// @Summary 发布菜单
// @Description 发布菜单
// @Param data body MenuRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/menus [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	// 判断菜单是否重复
	var menuModelList []models.MenuModel
	count := global.DB.Find(&menuModelList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.FailWithMsg("菜单重复", c)
		return
	}
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	// 菜单入库
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		res.FailWithMsg("发布菜单失败", c)
		return
	}
	// 菜单没有图片的情况下
	if len(cr.ImageSortList) == 0 {
		res.OKWithMsg("菜单添加成功", c)
		return
	}
	var menuBannerList []models.MenuBannerModel
	// 菜单有图片的情况下
	for _, image := range cr.ImageSortList {
		// 判断图片是否真实存在
		count = global.DB.Take(&models.BannerModel{}, image.ImageID).RowsAffected
		if count == 0 {
			global.Log.Error(fmt.Sprintf("ID为%d的图片不存在", image.ImageID))
			continue
		}
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: image.ImageID,
			Sort:     image.Sort,
		})
	}
	// 中间表入库
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMsg("菜单图片关联失败", c)
		return
	}
	res.OKWithMsg("发布菜单成功", c)
}
