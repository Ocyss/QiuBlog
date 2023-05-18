package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"qiublog/db"
	"qiublog/model"
	"qiublog/utils/ask"
	"qiublog/utils/errmsg"
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

func AddMessage(c *gin.Context) (int, any) {
	var data struct {
		model.Message
		Token string `json:"token"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	ctx := context.Background()
	if db.Rdb.Get(ctx, "Captcha:Yes:"+data.Token).Err() != nil {
		return errmsg.ERROR_CAPTCHA, nil
	}
	db.Rdb.Del(ctx, "Captcha:Yes:"+data.Token)
	_ = tool.WxPush(fmt.Sprintf("QiuBlog提醒您:\n有一条新的留言\n\n昵称: %s\nQQ: %s\n邮箱: %s\n内容: %s", data.Name, data.Qq, data.Email, data.Content))
	return model.AddMessage(&data.Message), nil
}

func AddQuestion(c *gin.Context) (int, any) {
	var data struct {
		model.Question
		Token string `json:"token"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	ctx := context.Background()
	if db.Rdb.Get(ctx, "Captcha:Yes:"+data.Token).Err() != nil {
		return errmsg.ERROR_CAPTCHA, nil
	}
	db.Rdb.Del(ctx, "Captcha:Yes:"+data.Token)
	_ = tool.WxPush(fmt.Sprintf("QiuBlog提醒您:\n有一条新的问答\n\n昵称: %s\nQQ: %s\n邮箱: %s\n问题: %s", data.Name, data.Qq, data.Email, data.Question.Question))
	return model.AddQuestion(&data.Question), nil
}

func GetMessage(c *gin.Context) (int, any) {
	pageSize, pageNum := tool.PageTool(c) //分页最大数,分页偏移量
	_, admin := tool.IsAdmin(c, 100)
	data, total := model.GetMessage(pageSize, pageNum, admin)
	return errmsg.SUCCESS, gin.H{
		"list":  data,
		"total": total,
	}
}

func GetQuestion(c *gin.Context) (int, any) {
	pageSize, pageNum := tool.PageTool(c) //分页最大数,分页偏移量
	_, admin := tool.IsAdmin(c, 100)
	data, total := model.GetQuestion(pageSize, pageNum, admin)
	return errmsg.SUCCESS, gin.H{
		"list":  data,
		"total": total,
	}
}

func UpMessage(c *gin.Context) (int, any) {
	var data upData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	if !data.Show && data.Val == false {
		code = errmsg.ERROR_DE_APPROVE
	} else {
		code = model.UpMessage(data.Id, data.Val, data.Show, data.Message)
	}
	return code, nil
}

func DelMessage(c *gin.Context) (int, any) {
	var data delData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	return model.DelMessage(data.Id, data.Message), nil
}

func ReplyQuestion(c *gin.Context) (int, any) {
	var data reply
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	return model.ReplyQuestion(data.Id, data.Content), nil
}

func ClearMessage(c *gin.Context) (int, any) {
	var data clear
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	model.ClearMessage(data.Message, data.Question)
	return errmsg.SUCCESS, nil
}
