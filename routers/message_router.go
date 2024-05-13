package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (r RouterGroup) MessageRouter() {
	app := api.ApiGroupApp
	r.POST("/messages", middleware.JwtAuth(), app.MessageApi.MessageCreateView)
	r.GET("/messages", middleware.JwtAuth(), app.MessageApi.MessageListView)
	r.GET("/messages_all", middleware.JwtAuth(), app.MessageApi.MessageListAllView)
	r.POST("/messages_record", middleware.JwtAuth(), app.MessageApi.MessageRecordView)
	r.DELETE("/message_users", middleware.JwtAuth(), app.MessageApi.MessageRecordRemoveView)
}
