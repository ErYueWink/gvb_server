package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

type BannerModel struct {
	MODEL
	Path      string          `gorm:"comment:图片路径" json:"path"`         // 图片路径
	Hash      string          `gorm:"comment:图片的hash值" json:"hash"`     // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38;comment:图片名称" json:"name"` // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"imageType"`       // 文件类型 1：本地 2：七牛云
}

// BeforeDelete 删除之前触发的钩子函数
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	// 本地存储，删除数据库中的数据后还需要删除本地存储
	if b.ImageType == ctype.LOCAL {
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err.Error())
			return err
		}
	}
	return nil
}
