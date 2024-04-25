package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"log"
	"time"
)

// InitGorm 初始化数据库连接
func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		log.Println("未配置Host，取消Gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn() // 数据库连接地址
	var mysqlLogger logger.Interface
	// 日志级别为debug则显示所有日志，否则只显示错误日志
	if global.Config.Mysql.LogLevel == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Println("数据库连接失败")
		return nil
	}
	// 配置mysql相关信息
	sql, _ := db.DB()
	sql.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns) // 最大空闲连接数
	sql.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns) // 最多可容纳
	sql.SetConnMaxLifetime(time.Hour * 4)                 // 连接最大复用时间，不能超过mysql的wait_time
	return db
}
