package v1

import (
	"github.com/gin-gonic/gin"
	"qiublog/db"
	"qiublog/utils/errmsg"
)

func MainSetUV(c *gin.Context) {
	err := db.Rdb.PFAdd(db.Ctx, "main:uv", c.ClientIP()).Err()
	code = 200
	if err != nil {
		code = 500
		c.JSON(500, gin.H{"status": code,
			"message": errmsg.GetErrMsg(code)})
		return
	}
	c.JSON(200, gin.H{"status": code,
		"message": errmsg.GetErrMsg(code)})
}
