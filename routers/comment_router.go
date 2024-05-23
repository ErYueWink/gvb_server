package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (r RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	r.POST("/comments", middleware.JwtAuth(), app.CommentCreateView)
	r.GET("/comments/digg/:id", middleware.JwtAuth(), app.CommentDigg)
	r.GET("/comments/:id", app.CommentListView)
}
