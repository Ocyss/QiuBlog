package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"qiublog/db"
	"qiublog/model"
	"qiublog/utils/errmsg"
	"qiublog/utils/res"
)

// Register 注册
func Register(c *gin.Context) {
	var data struct {
		model.User
		Token string `json:"token"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		res.ErrParam(c)
		return
	}
	ctx := context.Background()
	if db.Rdb.Get(ctx, "Captcha:Yes:"+data.Token).Err() != nil {
		res.Err(c, errmsg.ERROR_CAPTCHA)
		return
	}
	db.Rdb.Del(ctx, "Captcha:Yes:"+data.Token)
	res.Return(c, model.Register(&data.User))
}

// Login 登陆
func Login(c *gin.Context) {
	var data struct {
		model.User
		Token string `json:"token"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		res.ErrParam(c)
		return
	}
	ctx := context.Background()
	if db.Rdb.Get(ctx, "Captcha:Yes:"+data.Token).Err() != nil {
		res.Err(c, errmsg.ERROR_CAPTCHA)
		return
	}
	code, uid, token := model.CheckLogin(&data.User)
	db.Rdb.Del(ctx, "Captcha:Yes:"+data.Token)
	res.ReturnData(c, code, gin.H{
		"uid":   uid,
		"token": token,
	})
}
