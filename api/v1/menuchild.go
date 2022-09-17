package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiublog/model"
	"qiublog/utils/errmsg"
)

// AddMenuchild 添加菜单子项
func AddMenuchild(c *gin.Context) {
	var data model.Menuchild
	err := c.ShouldBindJSON(&data)
	if err != nil {
		code = errmsg.ERROR_PARM_SO
	} else {
		code = model.AddMenu(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetMenuchild 获取菜单子项
func GetMenuchild(c *gin.Context) {
	data := model.GetMenu()
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    data,
	})
}

// SetMenuchild 设置菜单子项
func SetMenuchild(c *gin.Context) {
	var data []model.SetMenuChild
	var da []model.Menuchild
	err := c.ShouldBindJSON(&data)
	if err != nil {
		code = errmsg.ERROR_PARM_SO
	}
	code := model.SetMenu(data)
	if code == errmsg.SUCCESS {
		da = model.GetMenu()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    da,
	})
}
