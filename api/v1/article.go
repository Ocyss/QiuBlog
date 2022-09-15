package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiublog/model"
	"qiublog/utils/errmsg"
	"strconv"
)

var code int

// GetsArticle 获取文章列表
func GetsArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	tagid, _ := strconv.Atoi(c.Query("tagid"))
	menuid, _ := strconv.Atoi(c.Query("menuid"))
	if pageSize <= 0 && pageSize > 20 {
		pageSize = 10
	}
	if pageNum <= 0 {
		pageNum = -1
	}
	if tagid <= 0 {
		tagid = -1
	}
	if menuid <= 0 {
		menuid = -1
	}
	data := model.GetsArticle(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArticle 获取单文章
func GetArticle(c *gin.Context) {

}

// ReleaseArticle 发布文章
func ReleaseArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArticle(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// ModifyArticle 修改文章
func ModifyArticle(c *gin.Context) {

}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {

}
