package routers

import "gvb_server/api"

func (router RouterGroup) SettingsRouter() {
	app := api.ApiGroupApp.SettingsApi
	router.GET("/settings/site", app.SettingsSiteInfoView)
	router.PUT("/settings/site", app.SettingsSiteUpdateInfoView)
	router.GET("/settings/:name", app.SettingsInfoView)
	router.PUT("/settings/:name", app.SettingsInfoUpdateView)

}
