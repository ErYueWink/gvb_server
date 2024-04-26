package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type MenuNameResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// MenuNameListView 菜单名称列表
// @Tags 菜单管理
// @Summary 菜单名称列表
// @Description 菜单名称列表
// @Router /api/menu_names [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]MenuNameResponse}
func (MenuApi) MenuNameListView(c *gin.Context) {
	var menuNameResponse []MenuNameResponse
	err := global.DB.Find(&models.MenuModel{}).Select("id", "name", "path").Scan(&menuNameResponse).Error
	if err != nil {
		res.FailWithMsg("查询菜单名称列表失败", c)
		return
	}
	res.OKWithData(menuNameResponse, c)
}
