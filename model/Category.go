package model

import (
	"context"
	"encoding/json"
	"fmt"
	"qiublog/db"
	"qiublog/utils/errmsg"
	"time"
)

// Category 分类表
type Category struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	Name     string    `gorm:"type:varchar(255);not null;unique;comment:分类名" json:"name"`
	Mid      *uint     `gorm:"type:int;comment:菜单子项ID" json:"mid"`
	Homeshow bool      `gorm:"type:bool;comment:主页是否显示;default:true;not null;" json:"homeshow,omitempty"`
	Articles []Article `gorm:"foreignkey:Cid"`
}

// Menuchild 菜单子项表
type Menuchild struct {
	ID       uint       `gorm:"primarykey" json:"id"`
	Sort     uint       `gorm:"comment:排序字段" json:"sort"`
	Name     string     `gorm:"comment:菜单名;not null;unique" json:"name"`
	Ename    string     `gorm:"comment:英文名;not null;unique" json:"ename"`
	Logo     string     `gorm:"type:longtext;comment:图标名;" json:"logo"`
	Link     string     `gorm:"comment:路由名;not null;unique" json:"link"`
	ParentId uint       `gorm:"comment:父级id;not null" json:"parent_id"`
	Cids     []Category `gorm:"foreignkey:Mid"`
}

type SetMenuChild struct {
	Type     string `json:"type"`
	ID       uint   `json:"id,omitempty"`
	Sort     uint   `json:"sort,omitempty"`
	Name     string `json:"name,omitempty"`
	Ename    string `json:"ename,omitempty"`
	Logo     string `json:"logo,omitempty"`
	Link     string `json:"link,omitempty"`
	ParentId uint   `json:"parent_id"` //父级id
}

type SetCategory struct {
	Type     string `json:"type"`
	ID       uint   `json:"id"`
	Name     string `json:"name,omitempty"`
	Mid      uint   `json:"mid,omitempty"`
	Homeshow bool   `json:"homeshow,omitempty"`
}
type GetSingleMenuTy struct {
	ID    uint   `json:"id"`    //菜单id
	Name  string `json:"name"`  //菜单名
	Ename string `json:"ename"` //英文名
	Link  string `json:"link"`  //路由名
	Cids  []struct {
		ID       uint   `json:"id"`       //分类id
		Name     string `json:"name"`     //分类name
		Mid      uint   `json:"mid"`      //分类mid
		Homeshow bool   `json:"homeshow"` //分类首页是否显示
	} `gorm:"foreignkey:Mid" json:"cids"`
}

// SetCate  设置分类
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
	ctx := context.Background()
	db.Rdb.Del(ctx, "categorys")
	db.Rdb.Del(ctx, "menus")
	db.Rdb.Del(ctx, "menu")
	return errmsg.SUCCESS
}

// SetMenu 设置菜单子项
func SetMenu(dataL []SetMenuChild) int {
	tx := Db.Begin()
	for _, data := range dataL {
		menu := Menuchild{Sort: data.Sort, Name: data.Name, Ename: data.Ename, Logo: data.Logo, Link: data.Link}
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
			err := tx.Create(&menu).Error
			if err != nil {
				tx.Rollback()
				return errmsg.ERROR_SET_NEW
			}
		case "updata":
			err := tx.Model(&Menuchild{}).Where("id = ?", data.ID).Updates(menu).Error
			if err != nil {
				tx.Rollback()
				return errmsg.ERROR_SET_UP
			}
		default:
			return errmsg.ERROR_SET_TYPE
		}
	}
	tx.Commit()
	ctx := context.Background()
	db.Rdb.Del(ctx, "menus")
	db.Rdb.Del(ctx, "menu")
	return errmsg.SUCCESS
}

// AddMenu 添加菜单子项
func AddMenu(data *Menuchild) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	ctx := context.Background()
	db.Rdb.Del(ctx, "menus")
	db.Rdb.Del(ctx, "menu")
	return errmsg.SUCCESS
}

// GetMenu 获取菜单子项
func GetMenu() []Menuchild {
	var data []Menuchild
	ctx := context.Background()
	key := "menus"
	dataJson, err := db.Rdb.Get(ctx, key).Bytes()
	if err != nil {
		Db.Order("parent_id asc").Order("sort asc").Find(&data)
		dataJson, _ = json.Marshal(data)
		db.Rdb.Set(ctx, key, dataJson, 3*24*time.Hour)
	} else {
		_ = json.Unmarshal(dataJson, &data)
	}
	return data
}

// GetSingleMenu 获取单菜单项
func GetSingleMenu(link string) (int, GetSingleMenuTy) {
	var data GetSingleMenuTy
	ctx := context.Background()
	key := fmt.Sprintf("link:%s", link)
	dataJson, err := db.Rdb.HGet(ctx, "menu", key).Bytes()
	code := errmsg.SUCCESS
	if err != nil {
		err := Db.Model(Menuchild{}).Where("link=?", link).Find(&data).Error
		if err != nil {
			code = errmsg.ERROR
		} else {
			err = Db.Model(Category{}).Where("mid=?", data.ID).Find(&data.Cids).Error
			if err != nil {
				code = errmsg.ERROR
			}
		}
		dataJson, _ = json.Marshal(data)
		db.Rdb.HSet(ctx, "menu", key, dataJson)
	} else {
		_ = json.Unmarshal(dataJson, &data)
	}
	return code, data
}

// AddCategory 添加分类
func AddCategory(data *Category) (int, *uint) {
	if *data.Mid == 0 {
		data.Mid = nil
	}
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR, nil
	}
	ctx := context.Background()
	db.Rdb.Del(ctx, "categorys")
	return errmsg.SUCCESS, &data.ID
}

type GetCategoryTy struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Mid      uint   `json:"mid"`
	Homeshow bool   `json:"homeshow"`
}

// GetCategory 获取分类
func GetCategory(homeshow bool) []GetCategoryTy {
	var data []GetCategoryTy
	ctx := context.Background()
	key := fmt.Sprintf("show:%v", homeshow)
	dataJson, err := db.Rdb.HGet(ctx, "categorys", key).Bytes()
	if err != nil {
		Db.Model(&Category{}).
			Where(&Category{Homeshow: homeshow}).
			Find(&data)
		dataJson, err = json.Marshal(data)
		db.Rdb.HSet(ctx, "categorys", key, dataJson)
	} else {
		_ = json.Unmarshal(dataJson, &data)
	}
	return data
}

// GetMidCid 根据mid获取所有cid
func GetMidCid(mid int) []int {
	var data []Category
	var r []int
	ctx := context.Background()
	key := fmt.Sprintf("mid:%d", mid)
	rJson, err := db.Rdb.HGet(ctx, "categorys", key).Bytes()
	if err != nil {
		where := map[string]any{}
		if mid == 0 {
			where["homeshow"] = true
		} else if mid > 0 {
			where["mid"] = mid
		} else if mid == -1 {
			where["homeshow"] = false
		}
		err := Db.Where(where).Find(&data).Error
		if err != nil {
			return nil
		}
		for _, item := range data {
			r = append(r, int(item.ID))
		}
		rJson, _ = json.Marshal(r)
		db.Rdb.HSet(ctx, "categorys", key, rJson)
	} else {
		_ = json.Unmarshal(rJson, &r)
	}
	return r
}

// ModifyCategorys 批量修改分类
func ModifyCategorys(data *[]Category) int {
	tx := Db.Begin()
	for _, item := range *data {
		err = tx.Table("category").Select("*").Updates(item).Error
		if err != nil {
			tx.Rollback()
			return errmsg.ERROR
		}
	}
	tx.Commit()
	ctx := context.Background()
	db.Rdb.Del(ctx, "categorys")
	return errmsg.SUCCESS
}
