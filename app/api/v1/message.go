package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"qiublog/db"
	"qiublog/model"
	"qiublog/utils/errmsg"
	"qiublog/utils/res"
	"qiublog/utils/tool"
)

type (
	upData struct {
		Id      uint `json:"id"`
		Val     bool `json:"val"`
		Show    bool `json:"show"`
		Message bool `json:"message"`
	}
	delData struct {
		Id      uint `json:"id"`
		Message bool `json:"message"`
	}
	reply struct {
		Id      uint   `json:"id"`
		Content string `json:"content"`
	}
	clear struct {
		Message  bool `json:"message"`
		Question bool `json:"question"`
	}
)

func AddMessage(c *gin.Context) {
	var data struct {
		model.Message
		Token string `json:"token"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		res.ErrParam(c)
		return
	}
	ctx := context.Background()
	if db.Rdb.Get(ctx, "Captcha:Yes:"+data.Token).Err() != nil {
		res.Err(c, errmsg.ERROR_CAPTCHA)
		return
	}
	db.Rdb.Del(ctx, "Captcha:Yes:"+data.Token)
	if err := tool.Send(
		"QiuBlogBot:您有一条新的留言",
		fmt.Sprintf("\n\n昵称: %s\nQQ: %s\n邮箱: %s\n内容: %s", data.Name, data.Qq, data.Email, data.Content),
	); err != nil {
		log.Error("SendErr:", err)
	}
	res.OKData(c, model.AddMessage(&data.Message))
}

func AddQuestion(c *gin.Context) {
	var data struct {
		model.Question
		Token string `json:"token"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		res.ErrParam(c)
		return
	}
	ctx := context.Background()
	if db.Rdb.Get(ctx, "Captcha:Yes:"+data.Token).Err() != nil {
		res.Err(c, errmsg.ERROR_CAPTCHA)
		return
	}
	db.Rdb.Del(ctx, "Captcha:Yes:"+data.Token)

	if err := tool.Send(
		"QiuBlogBot:您有一条新的问答",
		fmt.Sprintf("\n\n昵称: %s\nQQ: %s\n邮箱: %s\n问题: %s", data.Name, data.Qq, data.Email, data.Question.Question),
	); err != nil {
		log.Error("SendErr:", err)
	}
	res.Return(c, model.AddQuestion(&data.Question))
}

func GetMessage(c *gin.Context) {
	pageSize, pageNum := tool.PageTool(c) //分页最大数,分页偏移量
	_, admin := tool.IsAdmin(c, 100)
	data, total := model.GetMessage(pageSize, pageNum, admin)

	res.OKData(c, gin.H{
		"list":  data,
		"total": total,
	})
}

func GetQuestion(c *gin.Context) {
	pageSize, pageNum := tool.PageTool(c) //分页最大数,分页偏移量
	_, admin := tool.IsAdmin(c, 100)
	data, total := model.GetQuestion(pageSize, pageNum, admin)
	res.OKData(c, gin.H{
		"list":  data,
		"total": total,
	})
}

func UpMessage(c *gin.Context) {
	var data upData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		res.ErrParam(c)
		return
	}
	if !data.Show && data.Val == false {
		code = errmsg.ERROR_DE_APPROVE
	} else {
		code = model.UpMessage(data.Id, data.Val, data.Show, data.Message)
	}
	res.Return(c, code)
}

func DelMessage(c *gin.Context) {
	var data delData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		res.ErrParam(c)
		return
	}
	res.Return(c, model.DelMessage(data.Id, data.Message))
}

func ReplyQuestion(c *gin.Context) {
	var data reply
	err := c.ShouldBindJSON(&data)
	if err != nil {
		res.ErrParam(c)
		return
	}
	res.Return(c, model.ReplyQuestion(data.Id, data.Content))
}

func ClearMessage(c *gin.Context) {
	var data clear
	err := c.ShouldBindJSON(&data)
	if err != nil {
		res.ErrParam(c)
		return
	}
	model.ClearMessage(data.Message, data.Question)
	res.OK(c)
}

func LikeMessage(c *gin.Context) {
	type d struct {
		Type string `json:"type"`
		Id   uint   `json:"id"`
	}
	var data d
	err := c.ShouldBindJSON(&data)
	if err != nil || (data.Type != "message" && data.Type != "question") {
		res.ErrParam(c)
		return
	}
	key := data.Type + ":like"
	ctx := context.Background()
	temp := db.Rdb.PFCount(ctx, key).Val()
	db.Rdb.PFAdd(ctx, key, c.ClientIP(), data.Id)
	if temp != db.Rdb.PFCount(ctx, key).Val() {
		err = model.LikeMessage(data.Type, data.Id)
		if err != nil {
			res.ErrData(c, errmsg.ERROR, err)
		} else {
			res.OK(c)
		}
	} else {
		res.OKData(c, "你已经点过赞啦!")
	}
}
