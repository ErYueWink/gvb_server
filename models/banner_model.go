package models

import "gvb_server/models/ctype"

type BannerModel struct {
	MODEL
	Path      string          `gorm:"comment:图片路径" json:"path"`         // 图片路径
	Hash      string          `gorm:"comment:图片的hash值" json:"hash"`     // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38;comment:图片名称" json:"name"` // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"imageType"`       // 文件类型 1：本地 2：七牛云
}
