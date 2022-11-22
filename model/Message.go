package model

import "qiublog/utils/errmsg"

type (
	Message struct {
		Model
		Name    string `gorm:"comment:昵称;not null" json:"name"`    //昵称
		Qq      string `gorm:"comment:QQ" json:"qq"`               //QQ
		Email   string `gorm:"comment:邮箱" json:"email"`            //邮箱
		Content string `gorm:"comment:内容;not null" json:"content"` //内容
		Like    int    `gorm:"comment:点赞;not null" json:"like"`    //点赞
		Check   bool   `gorm:"comment:审核状态;not null" json:"check"` //审核
		Show    bool   `gorm:"comment:显示;not null" json:"show"`    //显示
	}
	Question struct {
		Model
		Name     string `gorm:"comment:昵称" json:"name"`              //昵称
		Qq       string `gorm:"comment:QQ" json:"qq"`                //QQ
		Email    string `gorm:"comment:邮箱" json:"email"`             //邮箱
		Question string `gorm:"comment:问题;not null" json:"question"` //问题
		Reply    string `gorm:"comment:回答" json:"reply"`             //回答
		Like     int    `gorm:"comment:点赞;not null" json:"like"`     //点赞
		Check    bool   `gorm:"comment:审核状态;not null" json:"check"`  //审核
		Show     bool   `gorm:"comment:显示;not null" json:"show"`     //显示
	}
	MessageR struct {
		Model
		Name    string `json:"name"`    //昵称
		Qq      string `json:"qq"`      //QQ
		Email   string `json:"email"`   //邮箱
		Content string `json:"content"` //内容
		Like    int    `json:"like"`    //点赞
	}
	QuestionR struct {
		Model
		Name     string `json:"name"`     //昵称
		Qq       string `json:"qq"`       //QQ
		Email    string `json:"email"`    //邮箱
		Question string `json:"question"` //问题
		Reply    string `json:"reply"`    //回答
		Like     int    `json:"like"`     //点赞
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

func GetMessage(pageSize int, pageNum int, role int) (interface{}, int64) {
	var data []Message
	var datar []MessageR
	var total int64
	where := map[string]interface{}{}
	if role == -1 {
		where["show"] = true
		where["check"] = true
	}
	err = Db.
		Model(&Message{}).
		Where(where).
		Order("created_at desc").
		Limit(pageSize).
		Offset((pageNum - 1) * 10).
		Find(&data).
		Find(&datar).Error
	if err != nil {
		return nil, 0
	}
	Db.Model(&Message{}).Count(&total)
	if role == -1 {
		return &datar, total
	}
	return &data, total
}

func GetQuestion(pageSize int, pageNum int, role int) (interface{}, int64) {
	var data []Question
	var datar []QuestionR
	var total int64
	where := map[string]interface{}{}
	if role == -1 {
		where["show"] = true
		where["check"] = true
	}
	err = Db.
		Model(&Question{}).
		Where(where).
		Order("created_at desc").
		Limit(pageSize).
		Offset((pageNum - 1) * 10).
		Find(&data).Find(&datar).Error
	if err != nil {
		return nil, 0
	}
	Db.Model(&Question{}).Count(&total)
	if role == -1 {
		return &datar, total
	}
	return &data, total
}

func UpMessage(id uint, val bool, show bool, message bool) int {
	var ty string
	var table interface{}
	if show {
		ty = "show"
	} else {
		ty = "check"
	}
	if message {
		table = &Message{}
	} else {
		table = &Question{}
	}
	err = Db.Model(table).Where(id).Update(ty, val).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DelMessage(id uint, message bool) int {
	var table interface{}
	if message {
		table = &Message{}
	} else {
		table = &Question{}
	}
	err = Db.Delete(table, id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func ReplyQuestio(id uint, content string) int {
	err = Db.Model(&Question{}).Where(id).Update("reply", content).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
