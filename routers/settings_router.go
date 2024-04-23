package routers

import "gvb_server/api"

func (router RouterGroup) SettingsRouter() {
	app := api.ApiGroupApp.SettingsApi
	router.GET("/settings", app.SettingsInfoView)

}
