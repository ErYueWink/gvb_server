package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/images_api"
	"gvb_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi // 系统配置
	ImagesApi   images_api.ImagesApi     // 图片上传
	AdvertApi   advert_api.AdvertApi     // 广告管理
}

var ApiGroupApp = new(ApiGroup)
