package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"qiublog/db"
	"qiublog/model"
	"qiublog/utils/ask"
	"qiublog/utils/errmsg"
)

// Register 注册
func Register(c *gin.Context) (int, any) {
	var data struct {
		model.User
		Token string `json:"token"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	ctx := context.Background()
	if db.Rdb.Get(ctx, "Captcha:Yes:"+data.Token).Err() != nil {
		return errmsg.ERROR_CAPTCHA, nil
	}
	db.Rdb.Del(ctx, "Captcha:Yes:"+data.Token)
	return model.Register(&data.User), nil
}

// Login 登陆
func Login(c *gin.Context) (int, any) {
	var data struct {
		model.User
		Token string `json:"token"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}
	ctx := context.Background()
	if db.Rdb.Get(ctx, "Captcha:Yes:"+data.Token).Err() != nil {
		return errmsg.ERROR_CAPTCHA, nil
	}
	code, uid, token := model.CheckLogin(&data.User)
	db.Rdb.Del(ctx, "Captcha:Yes:"+data.Token)
	return code, gin.H{
		"uid":   uid,
		"token": token,
	}
}
