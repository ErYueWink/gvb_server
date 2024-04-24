package models

import "time"

// UserCollectModel 自定义第三张表 记录哪个用户收藏了哪篇文章
type UserCollectModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"comment:用户id" json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID string    `gorm:"size:32;comment:文章的es id"`
	CreatedAt time.Time `gorm:"comment:收藏的时间"`
}
