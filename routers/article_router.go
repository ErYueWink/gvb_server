package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (r RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp
	r.POST("/articles", middleware.JwtAuth(), app.ArticleApi.ArticleCreateView)
	r.GET("/articles", app.ArticleApi.ArticleListView)
	r.GET("/articles/:id", middleware.JwtAuth(), app.ArticleApi.ArticleDetailView)
	r.GET("/articles/detail", middleware.JwtAuth(), app.ArticleApi.ArticleDetailByTitleView)
	r.GET("/articles/calendar", middleware.JwtAuth(), app.ArticleApi.ArticleCalendarView)
	r.GET("/articles/collects", middleware.JwtAuth(), app.ArticleApi.ArticleCollCreateView)
}
