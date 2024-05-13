package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (r RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp
	r.POST("/articles", middleware.JwtAuth(), app.ArticleApi.ArticleCreateView)
}
