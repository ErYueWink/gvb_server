package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

// CommList 通用查询
func CommList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	// Debug模式开启日志显示
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	// 计算偏移量
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	// 页数默认为10
	if option.Limit == 0 {
		option.Limit = 10
	}
	// 获取总条数
	count = DB.Select("id").Find(&list).RowsAffected
	// 分页查询
	err = DB.Limit(option.Limit).Offset(offset).Find(&list).Error
	return list, count, err
}
