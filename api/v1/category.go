package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiublog/model"
	"qiublog/utils/ask"
	"qiublog/utils/errmsg"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.AddCategory(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCategory 获取分类
func GetCategory(c *gin.Context) {
	show, err := strconv.ParseBool(c.Query("show"))
	if err != nil {
		ask.ErrParam(c)
		return
	}
	data := model.GetCategory(show)

	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    data,
	})
}

// GetTags 获取全部标签
func GetTags(c *gin.Context) {
	data := model.GetTags()
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    data,
	})
}

// ModifyCategorys 批量修改分类
func ModifyCategorys(c *gin.Context) {
	var data []model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.ModifyCategorys(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
