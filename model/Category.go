package model

import (
	"gorm.io/gorm"
	"qiublog/utils/errmsg"
)

// Category 分类表
type Category struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(255);not null;unique;comment:分类名"`
}

// Menuchild 菜单子项表
type Menuchild struct {
	ID    uint   `gorm:"primarykey"`                                                        //菜单id
	Sort  uint   `gorm:"comment:排序字段;unique"`                                               //排序字段
	Name  string `gorm:"comment:菜单名;not null;unique" json:"name" binding:"required"`        //菜单名
	Ename string `gorm:"comment:英文名;not null;unique" json:"ename" binding:"required"`       //英文名
	Logo  string `gorm:"type:longtext;comment:图标名;not null" json:"logo" binding:"required"` //图标
	Link  string `gorm:"comment:路由名;not null;unique" json:"link" binding:"required"`        //路由名
}

// AddMenu 添加菜单子项
func AddMenu(data *Menuchild) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetMenu 获取菜单子项
func GetMenu() []Menuchild {
	var data []Menuchild
	err := db.Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return data
}

// AddCategory 添加分类
func AddCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
