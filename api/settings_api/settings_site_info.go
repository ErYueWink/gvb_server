package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/utils/res"
)

// SettingsSiteInfoView 获取网站信息
func (SettingsApi) SettingsSiteInfoView(c *gin.Context) {
	res.OKWithData(global.Config.SiteInfo, c)
}
