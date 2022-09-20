package model

import (
	"gorm.io/gorm"
	"qiublog/utils/errmsg"
)

// Category 分类表
type Category struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"type:varchar(255);not null;unique;comment:分类名" json:"name"`
	Menuchild Menuchild `gorm:"foreignKey:Mid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;comment:菜单子项"`
	Mid       uint      `gorm:"type:int;comment:菜单子项ID" json:"mid,omitempty"`
	Homeshow  bool      `gorm:"type:bool;comment:主页是否显示;default:true;not null;" json:"homeshow,omitempty"`
}

// Menuchild 菜单子项表
type Menuchild struct {
	ID    uint   `gorm:"primarykey"`                                     //菜单id
	Sort  uint   `gorm:"comment:排序字段"`                                   //排序字段
	Name  string `gorm:"comment:菜单名;not null;unique" json:"name"`        //菜单名
	Ename string `gorm:"comment:英文名;not null;unique" json:"ename"`       //英文名
	Logo  string `gorm:"type:longtext;comment:图标名;not null" json:"logo"` //图标
	Link  string `gorm:"comment:路由名;not null;unique" json:"link"`        //路由名
}

type SetMenuChild struct {
	Type  string `json:"type"`
	ID    uint   `json:"id,omitempty"`
	Sort  uint   `json:"sort,omitempty"`
	Name  string `json:"name,omitempty"`
	Ename string `json:"ename,omitempty"`
	Logo  string `json:"logo,omitempty"`
	Link  string `json:"link,omitempty"`
}

type SetCategory struct {
	Type     string `json:"type"`
	ID       uint   `json:"id"`
	Name     string `json:"name,omitempty"`
	Mid      uint   `json:"mid,omitempty"`
	Homeshow bool   `json:"homeshow,omitempty"`
}

// SetCate  设置菜单子项
func SetCate(dataL []SetCategory) int {
	tx := Db.Begin()
	for _, data := range dataL {
		switch data.Type {
		case "remove":
			err := tx.Delete(&Category{}, data.ID).Error
			if err != nil {
				tx.Rollback()
				return errmsg.ERROR_SET_RE
			}
		case "new":
			menu := Category{}
			err := tx.Create(&menu).Error
			if err != nil {
				tx.Rollback()
				return errmsg.ERROR_SET_NEW
			}
		case "updata":
		default:
			return errmsg.ERROR_SET_TYPE
		}
	}
	tx.Commit()
	return errmsg.SUCCESS
}

// SetMenu 设置菜单子项
func SetMenu(dataL []SetMenuChild) int {
	tx := Db.Begin()
	for _, data := range dataL {
		switch data.Type {
		case "sort":
			err := tx.Model(&Menuchild{}).Where("id = ?", data.ID).Update("sort", data.Sort).Error
			if err != nil {
				tx.Rollback()
				return errmsg.ERROR_SET_SO
			}
		case "remove":
			err := tx.Delete(&Menuchild{}, data.ID).Error
			if err != nil {
				tx.Rollback()
				return errmsg.ERROR_SET_RE
			}
		case "new":
			menu := Menuchild{Sort: data.Sort, Name: data.Name, Ename: data.Ename, Logo: data.Logo, Link: data.Link}
			err := tx.Create(&menu).Error
			if err != nil {
				tx.Rollback()
				return errmsg.ERROR_SET_NEW
			}
		case "updata":
			err := tx.Model(&Menuchild{}).Where("id = ?", data.ID).Updates(Menuchild{Sort: data.Sort, Name: data.Name, Ename: data.Ename, Logo: data.Logo, Link: data.Link}).Error
			if err != nil {
				tx.Rollback()
				return errmsg.ERROR_SET_UP
			}
		default:
			return errmsg.ERROR_SET_TYPE
		}
	}
	tx.Commit()

	return errmsg.SUCCESS
}

// AddMenu 添加菜单子项
func AddMenu(data *Menuchild) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetMenu 获取菜单子项
func GetMenu() []Menuchild {
	var data []Menuchild
	err := Db.Order("sort asc").Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return data
}

// AddCategory 添加分类
func AddCategory(data *Category) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCategory 获取分类
func GetCategory() []Category {
	var data []Category
	err := Db.Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return data
}

// GetMidCid 根据mid获取所有cid
func GetMidCid(mid int) []int {

}
