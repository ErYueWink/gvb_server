package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"gvb_server/api"
	"gvb_server/middleware"
)

var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (r RouterGroup) UserRouter() {
	app := api.ApiGroupApp
	r.Use(sessions.Sessions("sessionid", store))
	r.POST("/email_login", app.UserApi.EmailLoginView)
	r.GET("/users", middleware.JwtAuth(), app.UserApi.UserListView)
	r.PUT("/user_role", middleware.JwtAuth(), app.UserApi.UserUpdateRoleView)
	r.PUT("/user_password", middleware.JwtAuth(), app.UserApi.UserUpdatePasswordView)
	r.POST("/logout", middleware.JwtAuth(), app.UserApi.LogoutView)
	r.DELETE("/users", middleware.JwtAuth(), app.UserApi.UserRemoveView)
	r.POST("/user_bind_email", middleware.JwtAuth(), app.UserApi.UserBindEmailView)
}
