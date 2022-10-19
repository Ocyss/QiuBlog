package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"qiublog/model"
	"qiublog/utils/errmsg"
	"time"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	class := c.Request.FormValue("class")
	if err != nil {
		code = errmsg.ERROR_FILE_WRONG
		c.JSON(http.StatusOK, gin.H{
			"errno":   code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	fileSize := fileHeader.Size
	fileName := fileHeader.Filename
	if class == "Tags" || class == "User" || class == "Other" || class == "Article" {
		now := time.Now()
		uploadName := fmt.Sprintf("%s/%d%s", class, now.UnixNano(), path.Ext(fileName))
		url, code := model.UpLoadFile(uploadName, file, fileSize)
		if code == errmsg.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
				"errno":   0,
				"data": gin.H{
					"url": url,
				},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"errno":   code,
				"message": errmsg.GetErrMsg(code),
			})
		}
	} else {
		code := errmsg.ERROR_CLASS_WRONG
		c.JSON(http.StatusOK, gin.H{
			"errno":   code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}
