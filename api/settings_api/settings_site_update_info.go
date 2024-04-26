package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/res"
)

// SettingsSiteUpdateInfoView 修改网站信息
// @Tags 系统管理
// @Summary 编辑网站信息
// @Description 编辑网站信息
// @Param data body config.SiteInfo true "编辑网站信息的参数"
// @Param token header string  true  "token"
// @Router /api/settings/site [put]
// @Produce json
// @Success 200 {object} res.Response{data=config.SiteInfo}
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
