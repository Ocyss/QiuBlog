package ask

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiublog/utils/errmsg"
)

func ErrParam(c *gin.Context) {
	code := errmsg.ERROR_PARAM
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
