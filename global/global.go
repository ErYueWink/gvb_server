package global

import (
	"github.com/cc14514/go-geoip2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/config"
)

var (
	Config   *config.Config // 配置文件
	DB       *gorm.DB       // 数据库连接
	Log      *logrus.Logger
	MysqlLog logger.Interface
	AddrDB   *geoip2.DBReader
)
