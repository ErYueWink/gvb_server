package routers

import "gvb_server/api"

func (r RouterGroup) MenuRouter() {
	app := api.ApiGroupApp
	r.POST("/menus", app.MenuApi.MenuCreateView)
	r.GET("/menus", app.MenuApi.MenuListView)
	r.GET("/menus/:id", app.MenuApi.MenuDetailView)
	r.GET("/menus/detail", app.MenuApi.MenuDetailByPathView)
	r.DELETE("/menus", app.MenuApi.MenuRemoveView)
	r.PUT("/menus/:id", app.MenuApi.MenuUpdateView)
	r.GET("/menu_names", app.MenuApi.MenuNameListView)
}
