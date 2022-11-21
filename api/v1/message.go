package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiublog/model"
	"qiublog/utils/ask"
	"qiublog/utils/errmsg"
	"qiublog/utils/tool"
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
	data, total := model.GetMessage(pageSize, pageNum)
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
	data, total := model.GetQuestion(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    data,
		"total":   total,
	})
}
