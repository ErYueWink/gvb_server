package routers

import "gvb_server/api"

func (r RouterGroup) MenuRouter() {
	app := api.ApiGroupApp
	r.POST("/menus", app.MenuApi.MenuCreateView)
	r.GET("/menus", app.MenuApi.MenuListView)
}
