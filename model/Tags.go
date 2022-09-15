package model

// Tags 标签
type Tags struct {
	ID       uint     `gorm:"primarykey"`                                        // 标签ID
	Name     string   `gorm:"type:varchar(20);not null;comment:标签名" json:"name"` //标签名
	Logo     string   `gorm:"type:varchar(20);comment:LOGO" json:"logo"`         //LOGO
	Category Category `gorm:"foreignKey:Cid"`                                    //分类名
	Cid      int      `gorm:"type:int;comment:分类名id"  json:"menuid"`             //分类名ID
}

// ArticleTags 标签文章对应表
type ArticleTags struct {
	ID      uint    `gorm:"primarykey"`
	Article Article `gorm:"foreignKey:Aid"`
	Aid     int     `gorm:"type:int;comment:文章ID"` //文章ID
	Tags    Tags    `gorm:"foreignKey:Tid"`
	Tid     int     `gorm:"type:int;comment:标签ID"` //标签ID
}
