package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// MenuRemoveView 删除菜单
// @Tags 菜单管理
// @Summary 删除菜单
// @Description 删除菜单
// @Param data body models.RemoveRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/menus [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("图片不存在", c)
		return
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 清空关联数据
		err = global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}
		// 删除菜单数据
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		res.FailWithMsg("删除菜单失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("共删除%d个菜单", count), c)
}
