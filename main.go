package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	core.InitConf() // 初始化配置文件
	// 初始化日志
	global.Log = core.InitLogger()
	global.DB = core.InitGorm() // 初始化数据库连接
}
