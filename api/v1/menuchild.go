package v1

import (
	"github.com/gin-gonic/gin"
	"qiublog/model"
	"qiublog/utils/errmsg"
)

// AddMenuchild 添加菜单子项
func AddMenuchild(c *gin.Context) (int, any) {
	var data model.Menuchild
	err := c.ShouldBindJSON(&data)
	if err != nil {
		code = errmsg.ERROR_PARM_SO
	} else {
		code = model.AddMenu(&data)
	}
	return code, nil
}

// GetMenuchild 获取菜单子项
func GetMenuchild(c *gin.Context) (int, any) {
	return errmsg.SUCCESS, model.GetMenu()
}

// GetSingleMenuItem 获取单菜单项
func GetSingleMenuItem(c *gin.Context) (int, any) {
	MenuLink := c.Query("link")
	return model.GetSingleMenu(MenuLink)
}

// SetMenuchild 设置菜单子项
func SetMenuchild(c *gin.Context) (int, any) {
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
	return code, da
}
