package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func AddMessage(c *gin.Context) {
	var data model.Message
	err := c.ShouldBindJSON(&data)
	if err != nil {
		ask.ErrParam(c)
		return
	}
	code := model.AddMessage(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func AddQuestion(c *gin.Context) {
	var data model.Question
	err := c.ShouldBindJSON(&data)
	if err != nil {
		ask.ErrParam(c)
		return
	}
	code := model.AddQuestion(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetMessage(c *gin.Context) {
	pageSize, pageNum := tool.PageTool(c) //分页最大数,分页偏移量
	cRole, ok := c.Get("role")
	role, _ := cRole.(int)
	if !ok {
		role = -1
	}
	data, total := model.GetMessage(pageSize, pageNum, role)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    data,
		"total":   total,
	})
}

func GetQuestion(c *gin.Context) {
	pageSize, pageNum := tool.PageTool(c) //分页最大数,分页偏移量
	cRole, ok := c.Get("role")
	role, _ := cRole.(int)
	if !ok {
		role = -1
	}
	data, total := model.GetQuestion(pageSize, pageNum, role)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    data,
		"total":   total,
	})
}

func UpMessage(c *gin.Context) {
	var data upData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		ask.ErrParam(c)
		return
	}
	if !data.Show && data.Val == false {
		code = errmsg.ERROR_DE_APPROVE
	} else {
		code = model.UpMessage(data.Id, data.Val, data.Show, data.Message)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func DelMessage(c *gin.Context) {
	var data delData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		ask.ErrParam(c)
		return
	}
	code = model.DelMessage(data.Id, data.Message)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func ReplyQuestio(c *gin.Context) {
	var data reply
	err := c.ShouldBindJSON(&data)
	if err != nil {
		ask.ErrParam(c)
		return
	}
	code = model.ReplyQuestio(data.Id, data.Content)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
