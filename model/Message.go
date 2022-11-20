package model

import "qiublog/utils/errmsg"

type (
	Message struct {
		Model
		Name    string `gorm:"comment:昵称;not null" json:"name"`              //昵称
		Qq      string `gorm:"comment:QQ" json:"qq"`                         //QQ
		Email   string `gorm:"comment:邮箱" json:"email"`                      //邮箱
		Content string `gorm:"comment:内容;not null" json:"content"`           //内容
		Like    int    `gorm:"comment:点赞;not null" json:"like"`              //点赞
		Check   bool   `gorm:"comment:审核状态;not null" json:"check,omitempty"` //审核
		Show    bool   `gorm:"comment:显示;not null" json:"show,omitempty"`    //显示
	}
	Question struct {
		Model
		Name     string `gorm:"comment:昵称" json:"name"`                       //昵称
		Qq       string `gorm:"comment:QQ" json:"qq"`                         //QQ
		Email    string `gorm:"comment:邮箱" json:"email"`                      //邮箱
		Question string `gorm:"comment:问题;not null" json:"question"`          //问题
		Reply    string `gorm:"comment:回答" json:"reply"`                      //回答
		Like     int    `gorm:"comment:点赞;not null" json:"like"`              //点赞
		Check    bool   `gorm:"comment:审核状态;not null" json:"check,omitempty"` //审核
		Show     bool   `gorm:"comment:显示;not null" json:"show,omitempty"`    //显示
	}
	Comment struct{}
)

func AddMessage(data *Message) int {
	err = Db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func AddQuestion(data *Question) int {
	err = Db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func GetMessage(pageSize int, pageNum int) (*[]Message, int64) {
	var data []Message
	var total int64
	err = Db.
		Order("created_at desc").
		Limit(pageSize).
		Offset((pageNum - 1) * 10).
		Find(&data).Error
	if err != nil {
		return nil, 0
	}
	Db.Model(&Message{}).Count(&total)
	return &data, total
}

func GetQuestion(pageSize int, pageNum int) (*[]Question, int64) {
	var data []Question
	var total int64
	err = Db.
		Order("created_at desc").
		Limit(pageSize).
		Offset((pageNum - 1) * 10).
		Find(&data).Error
	if err != nil {
		return nil, 0
	}
	Db.Model(&Question{}).Count(&total)
	return &data, total
}
