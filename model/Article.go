package model

import (
	"gorm.io/gorm"
	"qiublog/utils/errmsg"
)

// Article 文章
type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(20);not null;comment:标题" json:"title"` //标题
	Img      string   `gorm:"type:varchar(128);comment:首图" json:"img"`           //首图
	Desc     string   `gorm:"type:varchar(20);comment:描述" json:"desc"`           //描述
	Content  string   `gorm:"type:longtext;comment:内容" json:"content"`           //内容
	Category Category `gorm:"foreignKey:Cid;comment:分类" `                        //分类
	Cid      int      `gorm:"type:int;comment:分类ID" json:"cid,omitempty"`        //分类ID
}

// CreateArticle 添加文章
func CreateArticle(tx *gorm.DB, data *Article) (int, uint) {
	err := tx.Create(&data).Error
	if err != nil {
		return errmsg.ERROR, 0
	}
	return errmsg.SUCCESS, data.ID
}

// GetsArticle 获取文章列表
func GetsArticle(pageSize int, pageNum int, cid int, tid int, mid int) []Article {
	var Articles []Article
	var menuids  []int
	if mid != 0{
		menuids=
	}
	filter := []string{"ID", "CreatedAt", "UpdatedAt", "title", "desc", "menuchild", "tags"}
	err := Db.Select(filter).Where(&Article{Cid: cid}).Order("createdAt desc").Limit(pageSize).Offset((pageNum - 1) * 10).Find(&Articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return Articles
}

// GetArticle 获取单个文章
func GetArticle(Aid int) {

}
