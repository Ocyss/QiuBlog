package v1

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/wenlng/go-captcha/captcha"
	"qiublog/db"
	"qiublog/utils/errmsg"
	"qiublog/utils/res"
	"qiublog/utils/tool"
	"time"
)

func GetCaptcha(c *gin.Context) {
	capt := captcha.GetCaptcha()
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		res.Err(c, errmsg.ERROR)
		return
	}
	ctx := context.Background()
	dotsByte, _ := json.Marshal(dots)
	db.Rdb.Set(ctx, "Captcha:"+key, dotsByte, 2*time.Minute)
	res.OKData(c, gin.H{"image_base64": b64, "thumb_base64": tb64, "captcha_key": key})
}

func CheckCaptcha(c *gin.Context) {
	var data struct {
		Key  string `json:"key"`
		Dots string `json:"dots"`
	}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		res.ErrParam(c)
		return
	}

	if v := tool.CheckCaptcha(data.Key, data.Dots); v != "" {
		res.OKData(c, v)
		return
	}
	res.Err(c, errmsg.ERROR_CAPTCHA)
}
