package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/utils/res"
)

var (
	QQNAME    = "qq"
	JWTNAME   = "jwt"
	EMAILNAME = "email"
	QINIUNAME = "qiniu"
)

type SettingsUri struct {
	Name string `uri:"name" binding:"required"`
}

// SettingsInfoView 获取某一项的配置文件信息
// @Tags 系统管理
// @Summary 显示某一项的配置信息
// @Description 显示某一项的配置信息  site email qq qiniu jwt chat_group
// @Param name path string  true  "name"
// @Param token header string  true  "token"
// @Router /api/settings/{name} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	// 参数绑定失败
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	GetSettingInfo(cr.Name, c)
}

func GetSettingInfo(name string, c *gin.Context) {
	switch name {
	case QQNAME:
		res.OKWithData(global.Config.QQ, c)
	case JWTNAME:
		res.OKWithData(global.Config.Jwt, c)
	case EMAILNAME:
		res.OKWithData(global.Config.Email, c)
	case QINIUNAME:
		res.OKWithData(global.Config.QiNiu, c)
	default:
		res.FailWithMsg("没有该配置文件", c)
	}

}
