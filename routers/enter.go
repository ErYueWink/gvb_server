package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvb_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiRouterGroup := router.Group("/api")
	routerGroup := RouterGroup{apiRouterGroup}
	routerGroup.SettingsRouter() // 系统配置
	routerGroup.ImagesRouter()   // 文件上传配置
	routerGroup.AdvertRouter()   // 广告管理
	routerGroup.MenuRouter()     // 菜单管理
	routerGroup.UserRouter()     // 用户管理
	routerGroup.TagRouter()      // 标签管理
	routerGroup.MessageRouter()  // 消息管理
	routerGroup.ArticleRouter()  // 文章管理
	routerGroup.CommentRouter()  // 评论管理
	return router
}
