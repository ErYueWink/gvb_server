package main

import (
	"fmt"
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

// @title API文档
// @version 1.0
// @description 肖晓恋爱星球API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	core.InitConf() // 初始化配置文件
	// 初始化日志
	global.Log = core.InitLogger()
	global.DB = core.InitGorm() // 初始化数据库连接
	global.Redis = core.InitConnectRedis()
	global.EsClient = core.EsConnect()
	// 路由初始化
	addr := global.Config.System.Addr()
	router := routers.InitRouter()
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
	}
	global.Log.Info(fmt.Printf("项目运行在%s", addr))
	router.Run(addr)

}
