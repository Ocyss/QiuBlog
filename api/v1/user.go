package v1

import (
	"github.com/gin-gonic/gin"
	"qiublog/model"
	"qiublog/utils/ask"
)

// Register 注册
func Register(c *gin.Context) (int, any) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	return model.Register(&data), nil
}

// Login 登陆
func Login(c *gin.Context) (int, any) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	code, uid, token := model.CheckLogin(&data)
	c.SetCookie("token", token, 259200, "/", c.Request.Referer(), false, false)
	return code, gin.H{
		"uid": uid,
	}
}
