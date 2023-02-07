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
	"strconv"
)

var code int

// GetsArticle 获取文章列表
func GetsArticle(c *gin.Context) (int, any) {
	pageSize, pageNum := tool.PageTool(c)  //分页最大数,分页偏移量
	cid, _ := strconv.Atoi(c.Query("cid")) //分类ID
	mid, _ := strconv.Atoi(c.Query("mid")) //菜单ID
	tid, _ := strconv.Atoi(c.Query("tid")) //标签ID
	//cids := model.GetMidCid(mid)
	data, total := model.GetsArticle(pageSize, pageNum, cid, mid, tid)
	//统计每个分类和菜单的访问次数
	var miduv, ciduv, tiduv int64
	ctx := context.Background()
	{
		midUvKey := fmt.Sprintf("articles/uv/mid:%d;", mid)
		db.Rdb.PFAdd(ctx, midUvKey, c.ClientIP())
		miduv, _ = db.Rdb.PFCount(ctx, midUvKey).Result()
		cidUvKey := fmt.Sprintf("articles/uv/cid:%d;", cid)
		db.Rdb.PFAdd(ctx, cidUvKey, c.ClientIP())
		ciduv, _ = db.Rdb.PFCount(ctx, cidUvKey).Result()
		tidUvKey := fmt.Sprintf("articles/uv/tid:%d;", tid)
		db.Rdb.PFAdd(ctx, tidUvKey, c.ClientIP())
		tiduv, _ = db.Rdb.PFCount(ctx, tidUvKey).Result()
	}
	return errmsg.SUCCESS, gin.H{
		"data":  data,
		"total": total,
		"uv": gin.H{
			"cid_uv": ciduv,
			"mid_uv": miduv,
			"tid_uv": tiduv,
		},
	}
}

// GetArticle 获取单文章
func GetArticle(c *gin.Context) (int, any) {
	aid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ask.ErrParam()
	}
	var uv int64
	//统计每篇文章的浏览量
	{
		ctx := context.Background()
		articleUvKey := fmt.Sprintf("article/uv/aid:%d;", aid)
		db.Rdb.PFAdd(ctx, articleUvKey, c.ClientIP())
		uv, _ = db.Rdb.PFCount(ctx, articleUvKey).Result()
	}
	code, data := model.GetArticle(aid)
	return code, gin.H{
		"data": data,
		"uv":   uv,
	}
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
