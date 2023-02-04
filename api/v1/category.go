package v1

import (
	"github.com/gin-gonic/gin"
	"qiublog/model"
	"qiublog/utils/ask"
	"qiublog/utils/errmsg"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) (int, any) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code, id := model.AddCategory(&data)
	return code, gin.H{"id": id}
}

// GetCategory 获取分类
func GetCategory(c *gin.Context) (int, any) {
	show, err := strconv.ParseBool(c.Query("show"))
	if err != nil {
		return ask.ErrParam()
	}
	return errmsg.SUCCESS, model.GetCategory(show)
}

// GetTags 获取全部标签
func GetTags(c *gin.Context) (int, any) {
	return errmsg.SUCCESS, model.GetTags()
}

// ModifyCategorys 批量修改分类
func ModifyCategorys(c *gin.Context) (int, any) {
	var data []model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.ModifyCategorys(&data)
	return code, nil
}
