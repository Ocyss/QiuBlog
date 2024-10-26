package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiublog/utils"
	"qiublog/utils/errmsg"
)

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": errmsg.GetErrMsg(http.StatusOK),
	})
}
func OKData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": errmsg.GetErrMsg(http.StatusOK),
		"data":    data,
	})
}

func Return(c *gin.Context, code int) {
	ReturnData(c, code, nil)
}
func ReturnData(c *gin.Context, code int, data any) {
	if code == errmsg.SUCCESS {
		OKData(c, data)
	} else {
		ErrData(c, code, data)
	}
}

func Err(c *gin.Context, code int) {
	ErrData(c, code, nil)
}
func ErrData(c *gin.Context, code int, data any) {
	res := gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	}
	//判断是否debug模式，是的话返回错误信息
	if _, ok := data.(error); ok && utils.Debug {
		res["errmsg"] = data
	}
	c.AbortWithStatusJSON(http.StatusTooManyRequests, res)
}
func ErrParam(c *gin.Context) {
	Err(c, errmsg.ERROR_PARAM)
}
