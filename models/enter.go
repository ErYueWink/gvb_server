package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey;comment:id" json:"id,select($any)" structs:"-"` // 主键ID
	CreatedAt time.Time `gorm:"comment:创建时间" json:"created_at,select($any)" structs:"-"`  // 创建时间
	UpdatedAt time.Time `gorm:"comment:更新时间" json:"-" structs:"-"`                        // 更新时间
}

type PageInfo struct {
	Page  int    `json:"page" form:"page"`   // 页数
	Limit int    `json:"limit" form:"limit"` // 每页搜索条数
	Sort  string `json:"sort" form:"sort"`   // 排序字段
	Key   string `json:"key" form:"key"`     // 搜索关键字
}
