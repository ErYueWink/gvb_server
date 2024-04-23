package global

import (
	"gorm.io/gorm"
	"gvb_server/config"
)

var (
	Config *config.Config // 配置文件
	DB     *gorm.DB       // 数据库连接
)
