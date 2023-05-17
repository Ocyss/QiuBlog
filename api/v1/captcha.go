package v1

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/wenlng/go-captcha/captcha"
	"qiublog/db"
	"qiublog/utils/ask"
	"qiublog/utils/errmsg"
	"qiublog/utils/tool"
)

func GetCaptcha(_ *gin.Context) (int, any) {
	capt := captcha.GetCaptcha()
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		return errmsg.ERROR, err
	}
	ctx := context.Background()
	dotsByte, _ := json.Marshal(dots)
	db.Rdb.HSet(ctx, "Captcha", key, dotsByte)
	return errmsg.SUCCESS, gin.H{"image_base64": b64, "thumb_base64": tb64, "captcha_key": key}
}

func CheckCaptcha(c *gin.Context) (int, any) {
	var data struct {
		Key  string `json:"key"`
		Dots string `json:"dots"`
	}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		return ask.ErrParam()
	}

	if tool.CheckCaptcha(data.Key, data.Dots) {
		return errmsg.SUCCESS, nil
	}
	return errmsg.ERROR_CAPTCHA, nil
}
