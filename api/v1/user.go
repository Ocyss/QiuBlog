package v1

import (
	"github.com/gin-gonic/gin"
	"qiublog/model"
	"qiublog/utils/ask"
	"qiublog/utils/errmsg"
	"qiublog/utils/tool"
)

// Register 注册
func Register(c *gin.Context) (int, any) {
	var data struct {
		model.User
		Dot string `json:"dot"`
		Key string `json:"key"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	if !tool.CheckCaptcha(data.Key, data.Dot) {
		return errmsg.ERROR_CAPTCHA, nil
	}
	return model.Register(&data.User), nil
}

// Login 登陆
func Login(c *gin.Context) (int, any) {
	var data struct {
		model.User
		Dot string `json:"dot"`
		Key string `json:"key"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	if !tool.CheckCaptcha(data.Key, data.Dot) {
		return errmsg.ERROR_CAPTCHA, nil
	}
	code, uid, token := model.CheckLogin(&data.User)
	return code, gin.H{
		"uid":   uid,
		"token": token,
	}
}
