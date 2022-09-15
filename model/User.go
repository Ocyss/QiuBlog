package model

import "gorm.io/gorm"

// User 用户
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;comment:用户名" json:"username"` //用户名
	Password string `gorm:"type:varchar(20);not null;comment:密码" json:"password"`  //密码
	Relation string `gorm:"type:varchar(20);not null;comment:关系" json:"relation"`  //关系
	Role     int    `gorm:"type:int;comment:权限" json:"role"`                       //权限0管理员1朋友2好友3访客
	Avatar   string `gorm:"type:varchar(20);comment:头像" json:"avatar"`             //头像
}
