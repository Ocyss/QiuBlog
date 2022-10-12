package model

import (
	"gorm.io/gorm"
	"qiublog/utils/errmsg"
)

// Article 文章
type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(64);not null;comment:标题" json:"title"` //标题
	Img     string `gorm:"type:varchar(128);comment:首图" json:"img"`           //首图
	Desc    string `gorm:"type:varchar(255);comment:描述" json:"desc"`          //描述
	Content string `gorm:"type:longtext;comment:内容" json:"content"`           //内容
	Cid     *int   `gorm:"type:int;comment:分类ID;" json:"cid,omitempty"`       //分类ID
	Tags    []Tags `gorm:"many2many:article_tags" json:"tags"`
}

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

// GetsArticle 获取文章列表
func GetsArticle(pageSize int, pageNum int, cid *int, cids []int) ([]Article, int64) {
	var Articles []Article
	var total int64
	err := Db.
		Preload("Tags").
		Where(Article{Cid: cid}).
		Order("created_at desc").
		Limit(pageSize).
		Offset((pageNum - 1) * 10).
		Find(&Articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	Db.Model(&Article{}).Where(Article{Cid: cid}).Count(&total)
	return Articles, total
}

// GetArticle 获取单个文章
func GetArticle(Aid int) (int, *Article) {
	var data Article
	err = Db.Preload("Tags").Where("id=?", Aid).Find(&data).Error
	if err != nil {
		return errmsg.ERROR, nil
	}
	return errmsg.SUCCESS, &data
}
