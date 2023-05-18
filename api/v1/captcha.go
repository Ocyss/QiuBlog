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
	"time"
)

func GetCaptcha(_ *gin.Context) (int, any) {
	capt := captcha.GetCaptcha()
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		return errmsg.ERROR, err
	}
	ctx := context.Background()
	dotsByte, _ := json.Marshal(dots)
	db.Rdb.Set(ctx, "Captcha:"+key, dotsByte, 2*time.Minute)
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

	if v := tool.CheckCaptcha(data.Key, data.Dots); v != "" {
		return errmsg.SUCCESS, v
	}
	return errmsg.ERROR_CAPTCHA, nil
}
