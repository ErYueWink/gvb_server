package routers

import "gvb_server/api"

func (r RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdvertApi
	r.GET("/advert", app.AdvertListView)
	r.POST("/advert", app.AdvertCreateView)
	r.PUT("/advert/:id", app.AdvertUpdateView)
	r.DELETE("/advert", app.AdvertRemoveView)
}
