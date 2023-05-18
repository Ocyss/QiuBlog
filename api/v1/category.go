package v1

import (
	"github.com/gin-gonic/gin"
	"qiublog/model"
	"qiublog/utils/res"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code, id := model.AddCategory(&data)
	res.ReturnData(c, code, id)
}

// GetCategory 获取分类
func GetCategory(c *gin.Context) {
	show, err := strconv.ParseBool(c.Query("show"))
	if err != nil {
		res.ErrParam(c)
		return
	}
	res.OKData(c, model.GetCategory(show))
}

// GetTags 获取全部标签
func GetTags(c *gin.Context) {
	res.OKData(c, model.GetTags())
}

// ModifyCategorys 批量修改分类
func ModifyCategorys(c *gin.Context) {
	var data []model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.ModifyCategorys(&data)
	res.Return(c, code)
}
