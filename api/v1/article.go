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
	cid, _ := strconv.Atoi(c.Query("cid")) //分类ID
	tid, _ := strconv.Atoi(c.Query("tid")) //标签ID
	mid, _ := strconv.Atoi(c.Query("mid")) //菜单子项ID
	if pageSize <= 0 && pageSize > 20 {
		pageSize = 10
	}
	if pageNum <= 0 {
		pageNum = -1
	}
	data := model.GetsArticle(pageSize, pageNum, cid, tid, mid)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": gin.H{
			"articles": data,
		},
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArticle 获取单文章
func GetArticle(c *gin.Context) {

}

// ReleaseArticle 发布文章
func ReleaseArticle(c *gin.Context) {
	var data struct {
		model.Article
		Tags []*string `json:"tags"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		code = errmsg.ERROR_JSON
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		tx := model.Db.Begin()
		code, Aid := model.CreateArticle(
			tx, &model.Article{
				Cid:     data.Cid,
				Desc:    data.Desc,
				Title:   data.Title,
				Content: data.Content,
				Img:     data.Img,
			})
		if code == errmsg.ERROR {
			tx.Rollback()
		} else {
			code = model.AddsTags(tx, data.Tags, Aid)
			if code == errmsg.ERROR {
				tx.Rollback()
			}
		}
		if code == errmsg.SUCCESS {
			tx.Commit()
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})

	}

}

// ModifyArticle 修改文章
func ModifyArticle(c *gin.Context) {

}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {

}
