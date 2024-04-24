package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/res"
)

// SettingsSiteUpdateInfoView 修改网站信息
func (SettingsApi) SettingsSiteUpdateInfoView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = cr
	err = core.SetYaml() // 修改配置文件
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		global.Log.Error(err) // 修改配置文件失败
		return
	}
	res.OKWith(c)
}
