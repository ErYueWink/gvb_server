package routers

import "gvb_server/api"

func (r RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	r.POST("/images", app.ImagesUploadView)
	r.GET("/images", app.ImagesListView)
	r.PUT("/images", app.ImagesUpdateView)
	r.DELETE("/images", app.ImagesRemoveApi)
	r.GET("/images/name", app.ImagesNameListView)
	r.POST("/images/data", app.ImagesUploadDataView)
}
