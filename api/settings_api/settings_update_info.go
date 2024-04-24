package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/res"
)

// SettingsInfoUpdateView 修改某一项的配置文件信息
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	err = UpdateSettingInfo(cr.Name, c)
	if err != nil {
		res.FailWithMsg("修改配置文件失败", c)
		return
	}
	res.OKWithMsg("修改配置文件成功", c)
}

func UpdateSettingInfo(name string, c *gin.Context) error {
	switch name {
	case QQNAME:
		var qq config.QQ
		err := c.ShouldBindJSON(&qq)
		if err != nil {
			return err
		}
		global.Config.QQ = qq
		break
	case JWTNAME:
		var jwt config.Jwt
		err := c.ShouldBindJSON(&jwt)
		if err != nil {
			return err
		}
		global.Config.Jwt = jwt
		break
	case EMAILNAME:
		var email config.Email
		err := c.ShouldBindJSON(&email)
		if err != nil {
			return err
		}
		global.Config.Email = email
		break
	case QINIUNAME:
		var qiniu config.QiNiu
		err := c.ShouldBindJSON(&qiniu)
		if err != nil {
			return err
		}
		global.Config.QiNiu = qiniu
		break
	default:
		res.FailWithMsg("没有该配置文件", c)
	}
	err := core.SetYaml()
	if err != nil {
		res.FailWithMsg("没有该配置文件", c)
		global.Log.Error("修改配置文件失败")
		return err
	}
	return nil
}
