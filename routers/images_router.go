package routers

import "gvb_server/api"

func (r RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	r.POST("/images", app.ImagesUploadView)
	r.GET("/images", app.ImagesListView)
}
