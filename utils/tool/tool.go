package tool

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// SplitToIntList 字符串切割成int数组
func SplitToIntList(str string, sep string) (intList []int) {
	if str == "" {
		return
	}
	strList := strings.Split(str, sep)
	if len(strList) == 0 {
		return
	}
	for _, item := range strList {
		if item == "" {
			continue
		}
		val, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			// logs.CtxError(ctx, "ParseInt fail, err=%v, str=%v, sep=%v", err, str, sep) // 此处打印出log报错信息
			continue
		}
		intList = append(intList, int(val))
	}
	return
}

// PageTool 分页通用获取
func PageTool(c *gin.Context) (int, int) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize")) //分页最大数
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))   //分页偏移量

	if pageSize <= 0 || pageSize > 20 {
		pageSize = 6
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	return pageSize, pageNum
}
