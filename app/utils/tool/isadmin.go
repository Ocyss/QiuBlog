package tool

import (
	"github.com/gin-gonic/gin"
	"qiublog/utils/errmsg"
	"qiublog/utils/middleware"
	"time"
)

func IsAdmin(c *gin.Context, role int) (*middleware.MyClaims, bool) {
	ckToken := c.GetHeader("token")
	key, tCode := middleware.CheckToken(ckToken)
	if tCode == errmsg.ERROR || key.Role < role || time.Now().Unix() > key.ExpiresAt {
		return nil, false
	}
	return key, true
}
