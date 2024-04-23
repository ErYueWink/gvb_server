package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	core.InitConf() // 初始化配置文件
	// 初始化日志
	global.Log = core.InitLogger()
	global.DB = core.InitGorm() // 初始化数据库连接
	// 路由初始化
	addr := global.Config.System.Addr()
	router := routers.InitRouter()
	global.Log.Info(fmt.Printf("项目运行在%s", addr))
	router.Run(addr)

}
