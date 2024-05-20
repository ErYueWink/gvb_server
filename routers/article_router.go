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
	r.GET("/articles/collects", middleware.JwtAuth(), app.ArticleApi.ArticleCollListView)
	r.POST("/articles/collects", middleware.JwtAuth(), app.ArticleApi.ArticleCollCreateView)
	r.DELETE("/articles/collects", middleware.JwtAuth(), app.ArticleApi.ArticleCollBatchRemoveView)
	r.GET("/articles/content/:id", middleware.JwtAuth(), app.ArticleApi.ArticleContentByIDView)
	r.POST("/articles/digg", middleware.JwtAuth(), app.ArticleApi.ArticleDiggView)
	r.GET("/article_id_title", middleware.JwtAuth(), app.ArticleApi.ArticleIDTitleListView)
	r.GET("/categories", middleware.JwtAuth(), app.ArticleApi.ArticleCategoryListView)
	r.GET("/articles/text", app.ArticleApi.FullTextSearchView)
}
