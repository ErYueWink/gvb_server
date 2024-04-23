package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	core.InitConf()             // 初始化配置文件
	global.DB = core.InitGorm() // 初始化数据库连接
	fmt.Println(global.DB)
}
