package model

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"qiublog/db"
	"qiublog/utils/errmsg"
	"strconv"
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
		Uv        int64     `json:"uv"`
	}
	Coding struct {
		Ids   []string
		Total int64
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
	ctx := context.Background()
	//删除所有缓存索引,保证数据一致
	db.Rdb.Del(ctx, "articles")
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
	ctx := context.Background()
	//保证数据一致性，直接删除缓存数据，下次自动从数据库获取最新数据
	db.Rdb.HDel(ctx, "article", strconv.Itoa(id))
	return errmsg.SUCCESS
}

// GetsArticle 获取文章列表
func GetsArticle(pageSize int, pageNum int, cid int, mid int, tid int) ([]Articles, int64) {
	// 存在缓存中的json数据
	var dataJson []byte
	//缓存中的格式
	var data Coding
	//需要返回的Articles
	res := make([]Articles, 0)

	ctx := context.Background()
	//根据key和条件，存进缓存
	articlesHset := func(key string, where map[string]interface{}) {
		Db.
			Model(Article{}).
			Where(where).
			Order("created_at desc").
			Pluck("id", &data.Ids).
			Count(&data.Total)
		dataJson, _ = json.Marshal(data)
		db.Rdb.HSet(ctx, "articles", key, dataJson)
	}
	//根据分类或者菜单获取文章列表
	if cid != 0 {
		cidKey := fmt.Sprintf("cid:%d", cid)
		dataJson, err = db.Rdb.HGet(ctx, "articles", cidKey).Bytes()
		if err != nil {
			articlesHset(cidKey, map[string]interface{}{"cid": cid})
		}
	} else if tid != 0 {
		tidKey := fmt.Sprintf("tid:%d", tid)
		dataJson, err = db.Rdb.HGet(ctx, "articles", tidKey).Bytes()
		if err != nil {
			var tdata Tags
			Db.Preload("Article.Tags").Take(&tdata, tid)
			for _, v := range tdata.Article {
				data.Ids = append(data.Ids, strconv.Itoa(int(v.ID)))
			}
			dataJson, _ = json.Marshal(data)
			db.Rdb.HSet(ctx, "articles", tidKey, dataJson)
			fmt.Println(tidKey, data)
		}
	} else {
		midKey := fmt.Sprintf("mid:%d", mid)
		dataJson, err = db.Rdb.HGet(ctx, "articles", midKey).Bytes()
		if err != nil {
			if mid == -2 {
				articlesHset(midKey, nil)
			} else {
				articlesHset(midKey, map[string]interface{}{"cid": GetMidCid(mid)})
			}
		}
	}

	_ = json.Unmarshal(dataJson, &data)
	//防止越界
	var aids []string
	if pageNum > len(data.Ids) {
		aids = nil
	} else if pageNum == -1 {
		aids = data.Ids
	} else if pageNum+pageSize > len(data.Ids) {
		aids = data.Ids[pageNum-1:]
	} else {
		aids = data.Ids[pageNum-1 : pageNum-1+pageSize]
	}
	for _, v := range aids {
		var d Articles
		var dd Article
		articleJson, err := db.Rdb.HGet(ctx, "article", v).Bytes()
		articleUv, _ := db.Rdb.PFCount(ctx, fmt.Sprintf("article/uv/aid:%s;", v)).Result()
		if err != nil {
			Db.Preload("Tags").Take(&dd, v)
			articleJson, _ = json.Marshal(dd)
			db.Rdb.HSet(ctx, "article", v, articleJson)
		} else {
			_ = json.Unmarshal(articleJson, &dd)
		}
		//优化，列表不返回文章内容
		d = Articles{dd.ID, dd.CreatedAt, dd.UpdatedAt, dd.Title, dd.Img, dd.Desc, dd.Cid, dd.Tags, articleUv}
		res = append(res, d)
	}

	return res, data.Total
}

// GetArticle 获取单个文章
func GetArticle(Aid int) (int, *Article) {
	var data Article
	ctx := context.Background()
	articleJson, err := db.Rdb.HGet(ctx, "article", strconv.Itoa(Aid)).Bytes()
	if err != nil {
		err = Db.Preload("Tags").Where("id=?", Aid).Find(&data).Error
		if err != nil {
			return errmsg.ERROR, nil
		}
		articleJson, _ = json.Marshal(&data)
		db.Rdb.HSet(ctx, "article", strconv.Itoa(Aid), articleJson)
	} else {
		_ = json.Unmarshal(articleJson, &data)
	}
	return errmsg.SUCCESS, &data
}
