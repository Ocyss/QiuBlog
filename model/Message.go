package model

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"qiublog/db"
	"qiublog/utils/errmsg"
	"qiublog/utils/tool"
	"strconv"
)

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

func ClearMessage(message bool, question bool) {
	ctx := context.Background()
	if message {
		db.Rdb.HDel(ctx, "messages", "admin:true")
		db.Rdb.HDel(ctx, "messages", "admin:false")
		db.Rdb.Del(ctx, "message")
	}
	if question {
		db.Rdb.HDel(ctx, "questions", "admin:true")
		db.Rdb.HDel(ctx, "questions", "admin:false")
		db.Rdb.Del(ctx, "question")
	}
}

// AddMessage 添加留言
func AddMessage(data *Message) int {
	err = Db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	ctx := context.Background()
	db.Rdb.HDel(ctx, "messages", "admin:true")
	db.Rdb.HDel(ctx, "messages", "admin:false")
	return errmsg.SUCCESS
}

// AddQuestion 添加提问
func AddQuestion(data *Question) int {
	err = Db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	ctx := context.Background()
	db.Rdb.HDel(ctx, "questions", "admin:true")
	db.Rdb.HDel(ctx, "questions", "admin:false")
	return errmsg.SUCCESS
}

// GetMessage 获取留言
func GetMessage(pageSize int, pageNum int, admin bool) ([]any, int64) {
	var data Coding
	var dataJson []byte
	res := make([]any, 0)
	ctx := context.Background()
	key := fmt.Sprintf("admin:%v", admin)
	dataJson, err := db.Rdb.HGet(ctx, "messages", key).Bytes()
	if err != nil {
		where := map[string]interface{}{}
		if !admin {
			where["show"] = true
			where["check"] = true
		}
		Db.
			Model(&Message{}).
			Where(where).
			Order("created_at desc").
			Pluck("id", &data.Ids).
			Count(&data.Total)
		dataJson, _ = json.Marshal(data)
		db.Rdb.HSet(ctx, "messages", key, dataJson)
	} else {
		_ = json.Unmarshal(dataJson, &data)
	}
	ids := tool.PageIds(pageNum, pageSize, data.Ids)
	for _, v := range ids {
		var m Message
		mJson, err := db.Rdb.HGet(ctx, "message", v).Bytes()
		if err != nil {
			Db.Take(&m, v)
			mJson, _ = json.Marshal(m)
			db.Rdb.HSet(ctx, "message", v, mJson)
		} else {
			_ = json.Unmarshal(mJson, &m)
		}
		if admin {
			res = append(res, m)
		} else {
			res = append(res, MessageR{m.Model, m.Name, m.Qq, m.Email, m.Content, m.Like})
		}
	}

	return res, data.Total
}

// GetQuestion 获取提问
func GetQuestion(pageSize int, pageNum int, admin bool) ([]any, int64) {
	var data Coding
	var dataJson []byte
	res := make([]any, 0)
	ctx := context.Background()
	key := fmt.Sprintf("admin:%v", admin)
	dataJson, err := db.Rdb.HGet(ctx, "questions", key).Bytes()
	if err != nil {
		where := map[string]interface{}{}
		if !admin {
			where["show"] = true
			where["check"] = true
		}
		Db.
			Model(&Question{}).
			Where(where).
			Order("created_at desc").
			Pluck("id", &data.Ids).
			Count(&data.Total)
		dataJson, _ = json.Marshal(data)
		db.Rdb.HSet(ctx, "questions", key, dataJson)
	} else {
		_ = json.Unmarshal(dataJson, &data)
	}
	ids := tool.PageIds(pageNum, pageSize, data.Ids)

	for _, v := range ids {
		var q Question
		mJson, err := db.Rdb.HGet(ctx, "question", v).Bytes()
		if err != nil {
			Db.Take(&q, v)
			mJson, _ = json.Marshal(q)
			db.Rdb.HSet(ctx, "question", v, mJson)
		} else {
			_ = json.Unmarshal(mJson, &q)
		}
		if admin {
			res = append(res, q)
		} else {
			res = append(res, QuestionR{q.Model, q.Name, q.Qq, q.Email, q.Question, q.Reply, q.Like})
		}
	}

	return res, data.Total
}

// UpMessage 更新留言
func UpMessage(id uint, val bool, show bool, message bool) int {
	var ty string
	var table interface{}
	var key string
	if show {
		ty = "show"
	} else {
		ty = "check"
	}
	ctx := context.Background()
	if message {
		key = "message"
		table = &Message{}
	} else {
		key = "question"
		table = &Question{}
	}
	err = Db.Model(table).Where(id).Update(ty, val).Error
	if err != nil {
		return errmsg.ERROR
	}
	db.Rdb.HDel(ctx, key+"s", "admin:true")
	db.Rdb.HDel(ctx, key+"s", "admin:false")
	db.Rdb.HDel(ctx, key, strconv.Itoa(int(id)))
	return errmsg.SUCCESS
}

// DelMessage 删除留言
func DelMessage(id uint, message bool) int {
	var table interface{}
	var key string
	ctx := context.Background()
	if message {
		key = "message"
		table = &Message{}
	} else {
		key = "question"
		table = &Question{}
	}
	err = Db.Delete(table, id).Error
	if err != nil {
		return errmsg.ERROR
	}
	db.Rdb.HDel(ctx, key+"s", "admin:true")
	db.Rdb.HDel(ctx, key+"s", "admin:false")
	db.Rdb.HDel(ctx, key, strconv.Itoa(int(id)))
	return errmsg.SUCCESS
}

// ReplyQuestion 回答
func ReplyQuestion(id uint, content string) int {
	err = Db.Model(&Question{}).Where(id).Update("reply", content).Error
	if err != nil {
		return errmsg.ERROR
	}
	ctx := context.Background()
	db.Rdb.HDel(ctx, "questions", "admin:true")
	db.Rdb.HDel(ctx, "questions", "admin:false")
	db.Rdb.HDel(ctx, "question", strconv.Itoa(int(id)))
	return errmsg.SUCCESS
}

// LikeMessage 点赞
func LikeMessage(t string, id uint) error {
	err := Db.Table(t).Where("id = ?", id).UpdateColumn("like", gorm.Expr("`like` + ?", 1)).Error
	if err != nil {
		return err
	}
	ctx := context.Background()
	db.Rdb.HDel(ctx, t, strconv.Itoa(int(id)))
	return nil
}
