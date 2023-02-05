package v1

import (
	"github.com/gin-gonic/gin"
	"qiublog/model"
	"qiublog/utils/ask"
	"qiublog/utils/errmsg"
	"qiublog/utils/tool"
	"strconv"
)

var code int

// GetsArticle 获取文章列表
func GetsArticle(c *gin.Context) (int, any) {
	pageSize, pageNum := tool.PageTool(c)  //分页最大数,分页偏移量
	cid, _ := strconv.Atoi(c.Query("cid")) //分类ID
	mid, _ := strconv.Atoi(c.Query("mid")) //菜单ID
	cids := model.GetMidCid(mid)
	data, total := model.GetsArticle(pageSize, pageNum, cid, cids)
	return errmsg.SUCCESS, gin.H{
		"data":  data,
		"total": total,
	}
}

// GetArticle 获取单文章
func GetArticle(c *gin.Context) (int, any) {
	aid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ask.ErrParam()
	}
	return model.GetArticle(aid)
}

// ReleaseArticle 发布文章
func ReleaseArticle(c *gin.Context) (int, any) {
	var data model.Article
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	tx := model.Db.Begin()
	code, Aid := model.CreateArticle(tx, &data)
	if code == errmsg.ERROR {
		tx.Rollback()
	} else if code == errmsg.SUCCESS {
		tx.Commit()
	}
	return code, gin.H{
		"Aid": Aid,
	}
}

// ModifyArticle 修改文章
func ModifyArticle(c *gin.Context) (int, any) {
	var data model.Article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ask.ErrParam()
	}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	tx := model.Db.Begin()
	code = model.ModifyArticle(tx, id, &data)
	if code == errmsg.ERROR {
		tx.Rollback()
	} else if code == errmsg.SUCCESS {
		tx.Commit()
	}
	return code, nil
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) (int, any) {
	return 0, nil
}

// TagGetArticle  根据标签获取所有文章
func TagGetArticle(c *gin.Context) (int, any) {
	tagId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ask.ErrParam()
	}
	data, total := model.TagGetArticle(tagId)
	return errmsg.SUCCESS, gin.H{
		"data":  data,
		"total": total,
	}
}
