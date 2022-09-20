package model

import (
	"gorm.io/gorm"
	"qiublog/utils/errmsg"
)

// Tags 标签
type Tags struct {
	ID    uint   `gorm:"primarykey"`                                               // 标签ID
	Name  string `gorm:"type:varchar(20);not null;unique;comment:标签名" json:"name"` //标签名
	Logo  string `gorm:"type:varchar(20);comment:LOGO" json:"logo,omitempty"`      //LOGO
	Color string `gorm:"type:varchar(20);comment:颜色" json:"color,omitempty"`       //颜色
}

// ArticleTags 标签文章对应表
type ArticleTags struct {
	ID      uint    `gorm:"primarykey"`
	Article Article `gorm:"foreignKey:Aid"`
	Aid     uint    `gorm:"type:int;comment:文章ID"` //文章ID
	Tags    Tags    `gorm:"foreignKey:Tid"`
	Tid     uint    `gorm:"type:int;comment:标签ID"` //标签ID
}

// AddsTags 添加标签列表
func AddsTags(tx *gorm.DB, data []*string, Aid uint) int {
	for _, item := range data {
		//= Tags{Name: *item}
		//err := tx.Create(&d).Error
		var tag Tags
		d := tx.Where(Tags{Name: *item}).Attrs(Tags{Name: *item}).FirstOrCreate(&tag)
		if err := d.Error; err != nil {
			return errmsg.ERROR
		}
		err = tx.Create(&ArticleTags{Aid: Aid, Tid: tag.ID}).Error
		if err != nil {
			return errmsg.ERROR
		}
	}
	return errmsg.SUCCESS
}
