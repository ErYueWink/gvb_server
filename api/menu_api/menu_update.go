package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// MenuUpdateView 更新菜单
// @Tags 菜单管理
// @Summary 更新菜单
// @Description 更新菜单
// @Param data body MenuRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Param id path int  true  "id"
// @Router /api/menus/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	menuId := c.Param("id")
	// 查询菜单是否存在
	var menuModel models.MenuModel
	count := global.DB.Take(&menuModel, menuId).RowsAffected
	if count == 0 {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	// 清空第三张表数据
	global.DB.Model(&menuModel).Association("Banners").Clear()
	// 判断前端有没有传来图片
	if len(cr.ImageSortList) > 0 {
		var banners []models.MenuBannerModel
		// 有传来图片
		for _, image := range cr.ImageSortList {
			count = global.DB.Take(&models.BannerModel{}, image.ImageID).RowsAffected
			if count == 0 {
				global.Log.Error("图片不存在")
				continue
			}
			banners = append(banners, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: image.ImageID,
				Sort:     image.Sort,
			})
		}
		// 第三张表入库
		err = global.DB.Create(&banners).Error
		if err != nil {
			res.FailWithMsg("菜单图片关联失败", c)
			return
		}
	}
	// 结构体转Map修改菜单表
	menuMap := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(&menuMap).Error
	if err != nil {
		res.FailWithMsg("修改菜单失败", c)
		return
	}
	res.OKWithMsg("修改菜单成功", c)
}
