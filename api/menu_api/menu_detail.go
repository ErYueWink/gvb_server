package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// MenuDetailView 查询菜单详情
// @Tags 菜单管理
// @Summary 菜单详情
// @Description 菜单详情
// @Param id path int  true  "id"
// @Router /api/menus/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=MenuResponse}
func (MenuApi) MenuDetailView(c *gin.Context) {
	menuId := c.Param("id")
	var menuModel models.MenuModel
	count := global.DB.Take(&menuModel, menuId).RowsAffected
	if count == 0 {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	var menuBannerModels []models.MenuBannerModel
	// 查连接表
	err := global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBannerModels, "menu_id = ?", menuId).Error
	if err != nil {
		res.FailWithMsg("查询连接表失败", c)
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
