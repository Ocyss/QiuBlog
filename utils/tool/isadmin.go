package tool

import (
	"github.com/gin-gonic/gin"
	"qiublog/middleware"
	"qiublog/utils/errmsg"
)

func IsAdmin(c *gin.Context) (*middleware.MyClaims, bool) {
	ckToken, _ := c.Cookie("token")
	key, tCode := middleware.CheckToken(ckToken)
	if tCode == errmsg.ERROR {
		return nil, false
	}
	return key, true
}
