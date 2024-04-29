package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (r RouterGroup) UserRouter() {
	app := api.ApiGroupApp
	r.POST("/email_login", app.UserApi.EmailLoginView)
	r.GET("/users", middleware.JwtAuth(), app.UserApi.UserListView)
	r.PUT("/user_role", middleware.JwtAuth(), app.UserApi.UserUpdateRoleView)
}
