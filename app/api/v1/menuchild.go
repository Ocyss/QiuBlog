package v1

import (
	"github.com/gin-gonic/gin"
	"qiublog/model"
	"qiublog/utils/errmsg"
	"qiublog/utils/res"
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
	res.Return(c, code)
}

// GetMenuchild 获取菜单子项
func GetMenuchild(c *gin.Context) {
	res.OKData(c, model.GetMenu())
}

// GetSingleMenuItem 获取单菜单项
func GetSingleMenuItem(c *gin.Context) {
	MenuLink := c.Query("link")
	code, data := model.GetSingleMenu(MenuLink)
	res.ReturnData(c, code, data)
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
	res.ReturnData(c, code, da)
}
