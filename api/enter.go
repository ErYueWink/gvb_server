package api

import (
	"gvb_server/api/images_api"
	"gvb_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi // 系统配置
	ImagesApi   images_api.ImagesApi     // 图片上传
}

var ApiGroupApp = new(ApiGroup)
