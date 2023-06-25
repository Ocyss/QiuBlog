package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"qiublog/model"
	"qiublog/utils/errmsg"
	"qiublog/utils/res"
	"time"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	class := c.Request.FormValue("class")
	if err != nil {
		res.Err(c, errmsg.ERROR_FILE_WRONG)
		return
	}
	fileSize := fileHeader.Size
	fileName := fileHeader.Filename
	if class == "Tags" || class == "User" || class == "Other" || class == "Article" {
		now := time.Now()
		uploadName := fmt.Sprintf("%s/%d%s", class, now.UnixNano(), path.Ext(fileName))
		code, url := model.UpLoadFile(uploadName, file, fileSize)
		res.ReturnData(c, code, url)
		return
	}

	res.Err(c, errmsg.ERROR_CLASS_WRONG)
}
