package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (r RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	r.POST("/tags", middleware.JwtAuth(), app.TagCreateView)
	r.DELETE("/tags", middleware.JwtAuth(), app.TagRemoveView)
	r.PUT("/tags/:id", middleware.JwtAuth(), app.TagUpdateView)
	r.PUT("/tags", middleware.JwtAuth(), app.TagListView)

}
