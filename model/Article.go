package model

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"qiublog/db"
	"qiublog/utils/errmsg"
	"time"
)

type (
	Article struct {
		Model
		Title   string `gorm:"type:varchar(64);not null;comment:标题" json:"title"` //标题
		Img     string `gorm:"type:varchar(128);comment:首图" json:"img"`           //首图
		Desc    string `gorm:"type:varchar(255);comment:描述" json:"desc"`          //描述
		Content string `gorm:"type:longtext;comment:内容" json:"content"`           //内容
		Cid     *int   `gorm:"type:int;comment:分类ID;" json:"cid,omitempty"`       //分类ID
		Tags    []Tags `gorm:"many2many:article_tags" json:"tags"`
	}
	Articles struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Title     string    `json:"title"`
		Img       string    `json:"img"`
		Desc      string    `json:"desc"`
		Cid       *int      `json:"cid,omitempty"`
		Tags      []Tags    `json:"tags"`
	}
)

// CreateArticle 添加文章
func CreateArticle(tx *gorm.DB, data *Article) (int, uint) {
	for index, item := range data.Tags {
		var tags Tags
		err := tx.
			Where(Tags{Name: item.Name}).
			Attrs(Tags{Name: item.Name, Logo: item.Logo, Color: item.Color}).
			FirstOrInit(&tags).Error
		if err != nil {
			return errmsg.ERROR, 0
		}
		data.Tags[index] = tags
	}
	err := tx.Create(&data).Error
	if err != nil {
		return errmsg.ERROR, 0
	}
	return errmsg.SUCCESS, data.ID
}

// ModifyArticle 修改文章
func ModifyArticle(tx *gorm.DB, id int, data *Article) int {
	for index, item := range data.Tags {
		var tags Tags
		err := tx.
			Where(Tags{Name: item.Name}).
			Attrs(Tags{Name: item.Name, Logo: item.Logo, Color: item.Color}).
			FirstOrInit(&tags).Error
		if err != nil {
			return errmsg.ERROR
		}
		data.Tags[index] = tags
	}
	err := tx.Model(&Article{}).Where("id", id).Updates(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetsArticle 获取文章列表
func GetsArticle(pageSize int, pageNum int, cid int, cids []int) ([]Articles, int64) {
	var article []Article
	var articlesJson []byte
	r := struct {
		Data  []Articles `json:"data"`
		Total int64      `json:"total"`
	}{}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	RedisKey := fmt.Sprintf("articles;pageSize:%d;pageNum:%d;cid:%d;cids:%s;", pageSize, pageNum, cid, fmt.Sprint(cids))
	articlesJson, err = db.Rdb.Get(ctx, RedisKey).Bytes()
	if err != nil {
		where := map[string]interface{}{}
		if cid != 0 {
			where["cid"] = cid
		} else if cids != nil {
			where["cid"] = cids
		}
		err = Db.
			Preload("Tags").
			Where(where).
			Order("created_at desc").
			Limit(pageSize).
			Offset((pageNum - 1) * pageSize).
			Find(&article).Error
		if err == nil || err != gorm.ErrRecordNotFound {
			Db.Model(&Article{}).Where(where).Count(&r.Total)
			for _, v := range article {
				r.Data = append(r.Data, Articles{v.ID, v.CreatedAt, v.UpdatedAt, v.Title, v.Img, v.Desc, v.Cid, v.Tags})
			}
		}
		articlesJson, _ = json.Marshal(r)
		db.Rdb.Set(ctx, RedisKey, articlesJson, 3*24*time.Hour) // 存3天
	}
	_ = json.Unmarshal(articlesJson, &r)
	return r.Data, r.Total
}

// GetArticle 获取单个文章
func GetArticle(Aid int) (int, *Article) {
	var data Article
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	redisKey := fmt.Sprintf("article;aid:%d;", Aid)
	articleJson, err := db.Rdb.Get(ctx, redisKey).Bytes()
	if err != nil {
		err = Db.Preload("Tags").Where("id=?", Aid).Find(&data).Error
		if err != nil {
			return errmsg.ERROR, nil
		}
		articleJson, _ = json.Marshal(&data)
		db.Rdb.Set(ctx, redisKey, articleJson, 3*24*time.Hour)
	} else {
		_ = json.Unmarshal(articleJson, &data)
	}
	return errmsg.SUCCESS, &data
}

// TagGetArticle  根据标签获取所有文章
func TagGetArticle(tagId int) (*Tags, int64) {
	var tag Tags
	var total int64
	err = Db.Preload("Article.Tags").Find(&tag, tagId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	total = Db.Model(&Tags{}).Association("Article").Count()
	return &tag, total
}
