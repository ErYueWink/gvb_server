package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/article_api"
	"gvb_server/api/images_api"
	"gvb_server/api/menu_api"
	"gvb_server/api/message_api"
	"gvb_server/api/settings_api"
	"gvb_server/api/tag_api"
	"gvb_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi // 系统配置
	ImagesApi   images_api.ImagesApi     // 图片上传
	AdvertApi   advert_api.AdvertApi     // 广告管理
	MenuApi     menu_api.MenuApi         // 菜单管理
	UserApi     user_api.UserApi         // 用户管理
	TagApi      tag_api.TagApi           // 标签管理
	MessageApi  message_api.MessageApi   // 消息管理
	ArticleApi  article_api.ArticleApi   // 文章管理
}

var ApiGroupApp = new(ApiGroup)
