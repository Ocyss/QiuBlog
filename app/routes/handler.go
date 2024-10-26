package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiublog/utils"
	"qiublog/utils/errmsg"
)

type MyHandler func(*gin.Context) (int, any)

func Handler() func(h MyHandler) gin.HandlerFunc {
	return func(h MyHandler) gin.HandlerFunc {
		return func(c *gin.Context) {
			code, data := h(c)
			req := gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			}
			if code == errmsg.SUCCESS {
				//判断数据类型
				if data != nil {
					req["data"] = data
				}
				c.JSON(errmsg.SUCCESS, req)
			} else {
				//判断是否debug模式，是的话返回错误信息
				if utils.Config.Server.AppMode == "development" {
					req["errmsg"] = data
				}
				c.AbortWithStatusJSON(http.StatusTooManyRequests, req)
			}
		}
	}
}
