package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"qiublog/model"
	"qiublog/utils/errmsg"
	"time"
)

func Upload(c *gin.Context) (int, any) {
	file, fileHeader, err := c.Request.FormFile("file")
	class := c.Request.FormValue("class")
	if err != nil {
		return errmsg.ERROR_FILE_WRONG, err
	}
	fileSize := fileHeader.Size
	fileName := fileHeader.Filename
	if class == "Tags" || class == "User" || class == "Other" || class == "Article" {
		now := time.Now()
		uploadName := fmt.Sprintf("%s/%d%s", class, now.UnixNano(), path.Ext(fileName))
		return model.UpLoadFile(uploadName, file, fileSize)
	} else {
		return errmsg.ERROR_CLASS_WRONG, nil
	}
}
