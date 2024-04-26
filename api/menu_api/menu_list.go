package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banner"`
}

// MenuListView 查询菜单列表
func (MenuApi) MenuListView(c *gin.Context) {
	var menuList []models.MenuModel
	var menuIDList []uint
	// 先查询菜单表
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ? ", menuIDList)
	menus := make([]MenuResponse, 0)
	for _, menu := range menuList {
		banners := make([]Banner, 0)
		for _, banner := range menuBanners {
			if menu.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: menu,
			Banners:   banners,
		})
	}
	res.OKWithList(menus, int64(len(menus)), c)
}
