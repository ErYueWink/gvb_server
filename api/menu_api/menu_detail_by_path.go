package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type MenuDetailRequest struct {
	Path string `form:"path"`
}

// MenuDetailByPathView 根据路径查询菜单详情
// @Tags 菜单管理
// @Summary 菜单详情,根据路径查
// @Description 菜单详情,根据路径查
// @Param data query MenuDetailRequest  true  "路径参数"
// @Router /api/menus/detail [get]
// @Produce json
// @Success 200 {object} res.Response{data=MenuResponse}
func (MenuApi) MenuDetailByPathView(c *gin.Context) {
	var cr MenuDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	// 根据路径查询菜单详情
	var menuModel models.MenuModel
	count := global.DB.Take(&menuModel, "path = ?", cr.Path).RowsAffected
	if count == 0 {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	var menuBannerModels []models.MenuBannerModel
	// 查询中间表数据
	err := global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBannerModels, "menu_id = ?", menuModel.ID).Error
	if err != nil {
		res.FailWithMsg("查询中间表数据失败", c)
		return
	}
	var banners = make([]Banner, 0)
	for _, model := range menuBannerModels {
		if menuModel.ID != model.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   model.BannerID,
			Path: model.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OKWithData(menuResponse, c)
}
