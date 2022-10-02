package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiublog/model"
	"qiublog/utils/ask"
	"qiublog/utils/errmsg"
	"qiublog/utils/tool"
	"strconv"
)

var code int

// GetsArticle 获取文章列表
func GetsArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cid, _ := strconv.Atoi(c.Query("cid")) //分类ID
	rcid := &cid
	cids := tool.SplitToIntList(c.Query("cids"), ",") //分类ID列表
	if pageSize <= 0 || pageSize > 20 {
		pageSize = 10
	}
	if pageNum <= 0 {
		pageNum = -1
	}
	if *rcid == 0 {
		rcid = nil
	}
	data, total := model.GetsArticle(pageSize, pageNum, rcid, cids)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": gin.H{
			"articles": data,
		},
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArticle 获取单文章
func GetArticle(c *gin.Context) {

}

//// ReleaseArticle 发布文章
//func ReleaseArticle(c *gin.Context) {
//	var data struct {
//		model.Article
//		Tags []*string `json:"tags"`
//	}
//	err := c.ShouldBindJSON(&data)
//	if err != nil {
//		ask.ErrParam(c)
//	} else {
//		tx := model.Db.Begin()
//		code, Aid := model.CreateArticle(
//			tx, &model.Article{
//				Cid:     data.Cid,
//				Desc:    data.Desc,
//				Title:   data.Title,
//				Content: data.Content,
//				Img:     data.Img,
//			})
//		if code == errmsg.ERROR {
//			tx.Rollback()
//		} else {
//			code = model.AddsTags(tx, data.Tags, Aid)
//			if code == errmsg.ERROR {
//				tx.Rollback()
//			}
//		}
//		if code == errmsg.SUCCESS {
//			tx.Commit()
//		}
//		c.JSON(http.StatusOK, gin.H{
//			"status":  code,
//			"message": errmsg.GetErrMsg(code),
//		})
//
//	}
//}

// ReleaseArticle 发布文章
func ReleaseArticle(c *gin.Context) {
	var data model.Article
	err := c.ShouldBindJSON(&data)
	if err != nil {
		ask.ErrParam(c)
		return
	}
	tx := model.Db.Begin()
	code, Aid := model.CreateArticle(tx, &data)
	if code == errmsg.ERROR {
		tx.Rollback()
	}
	if code == errmsg.SUCCESS {
		tx.Commit()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"Aid":     Aid,
	})
}

// ModifyArticle 修改文章
func ModifyArticle(c *gin.Context) {

}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {

}
