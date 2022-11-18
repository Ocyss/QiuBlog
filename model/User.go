package model

import (
	"gopkg.in/hlandau/passlib.v1"
	"qiublog/middleware"
	"qiublog/utils/errmsg"
)

// User 用户
type User struct {
	Model
	Username string `gorm:"type:varchar(40);not null;comment:用户名" json:"username"` //用户名
	Password string `gorm:"type:varchar(100);not null;comment:密码" json:"password"` //密码
	NickName string `gorm:"type:varchar(20)" json:"nickname"`                      //昵称
	Role     int    `gorm:"type:int;comment:权限" json:"role"`                       //权限0管理员1朋友2好友3访客
	Avatar   string `gorm:"type:varchar(20);comment:头像" json:"avatar"`             //头像
}

func CheckLogin(user *User) (int, uint, string) {
	var data User
	//根据用户名获取对应的全部数据
	err = Db.Where("username=?", user.Username).Find(&data).Error
	if err != nil {
		return errmsg.ERROR, 0, ""
	}
	//进行哈希值效验密码是否正确
	newHash, err := passlib.Verify(user.Password, data.Password)
	if err != nil {
		return errmsg.ERROR_PWDERR_WRONG, 0, ""
	}
	if newHash != "" {
		//登陆成功，判断是否需要更换哈希值
		Db.Model(&User{}).Where(data).Update("password", newHash)
	}
	token, tCode := middleware.SetToken(data.ID, data.Username, data.Role)
	if tCode != errmsg.SUCCESS {
		return tCode, 0, ""
	}
	return errmsg.SUCCESS, data.ID, token
}

func Register(user *User) int {
	hash, err := passlib.Hash(user.Password)
	if err != nil {
		// couldn't hash password for some reason
		return errmsg.ERROR
	}
	user.Password = hash
	err = Db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
