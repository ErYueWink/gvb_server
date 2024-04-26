package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/utils/res"
)

// SettingsSiteInfoView 获取网站信息
// @Tags 系统管理
// @Summary 显示网站信息
// @Description 显示网站信息
// @Router /api/settings/site [get]
// @Produce json
// @Success 200 {object} res.Response{data=config.SiteInfo}
func (SettingsApi) SettingsSiteInfoView(c *gin.Context) {
	res.OKWithData(global.Config.SiteInfo, c)
}
