package model

import (
	"context"
	"encoding/json"
	"qiublog/db"
	"time"
)

// Tags 标签
type (
	Tags struct {
		ID      uint      `gorm:"primarykey" json:"id,omitempty"`
		Name    string    `gorm:"type:varchar(20);not null;unique;comment:标签名" json:"name"` //标签名
		Logo    string    `gorm:"type:varchar(20);comment:LOGO" json:"logo,omitempty"`      //LOGO
		Color   string    `gorm:"type:varchar(20);comment:颜色" json:"color,omitempty"`       //颜色
		Article []Article `gorm:"many2many:article_tags"`
	}
	ReqTags struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
)

// GetTags 获取标签列表
func GetTags() *[]ReqTags {
	var data []ReqTags
	ctx := context.Background()
	key := "tags"
	dataJson, err := db.Rdb.Get(ctx, key).Bytes()
	if err != nil {
		Db.Model(Tags{}).Find(&data)
		dataJson, _ = json.Marshal(data)
		db.Rdb.Set(ctx, key, dataJson, 3*24*time.Hour)
	} else {
		_ = json.Unmarshal(dataJson, &data)
	}
	return &data
}

//// AddsTags 添加标签列表
//func AddsTags(tx *gorm.DB, data []*string, Aid uint) int {
//	for _, item := range data {
//		//= Tags{Name: *item}
//		//err := tx.Create(&d).Error
//		var tag Tags
//		d := tx.Where(Tags{Name: *item}).Attrs(Tags{Name: *item}).FirstOrCreate(&tag)
//		if err := d.Error; err != nil {
//			return errmsg.ERROR
//		}
//	}
//	return errmsg.SUCCESS
//}
