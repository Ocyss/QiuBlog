package v1

import (
	"github.com/gin-gonic/gin"
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
)

func AddMessage(c *gin.Context) (int, any) {
	var data model.Message
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	return model.AddMessage(&data), nil
}

func AddQuestion(c *gin.Context) (int, any) {
	var data model.Question
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	return model.AddQuestion(&data), nil
}

func GetMessage(c *gin.Context) (int, any) {
	pageSize, pageNum := tool.PageTool(c) //分页最大数,分页偏移量
	cRole, ok := c.Get("role")
	role, _ := cRole.(int)
	if !ok {
		role = -1
	}
	data, total := model.GetMessage(pageSize, pageNum, role)
	return errmsg.SUCCESS, gin.H{
		"data":  data,
		"total": total,
	}
}

func GetQuestion(c *gin.Context) (int, any) {
	pageSize, pageNum := tool.PageTool(c) //分页最大数,分页偏移量
	cRole, ok := c.Get("role")
	role, _ := cRole.(int)
	if !ok {
		role = -1
	}
	data, total := model.GetQuestion(pageSize, pageNum, role)
	return errmsg.SUCCESS, gin.H{
		"data":  data,
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

func ReplyQuestio(c *gin.Context) (int, any) {
	var data reply
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	return model.ReplyQuestio(data.Id, data.Content), nil
}
