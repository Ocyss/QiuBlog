package tool

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wenlng/go-captcha/captcha"
	"golang.org/x/net/context"
	"qiublog/db"
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
		pageSize = 10
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	return pageSize, pageNum
}

func PageIds(pageNum, pageSize int, dataIds []string) (ids []string) {
	n := len(dataIds)
	if pageNum > n || (pageNum-1)*pageSize > n {
		ids = nil
	} else if pageNum == -1 {
		ids = dataIds
	} else if pageNum*pageSize > n {
		ids = dataIds[(pageNum-1)*pageSize:]
	} else {
		ids = dataIds[(pageNum-1)*pageSize : pageNum*pageSize]
	}
	return
}

func CheckCaptcha(key, dots string) bool {
	if len(key) == 0 || len(dots) == 0 {
		return false
	}
	ctx := context.Background()
	dotsByte, err := db.Rdb.HGet(ctx, "Captcha", key).Bytes()
	if err != nil {
		return false
	}
	var (
		dct map[int]captcha.CharDot
		src []string
	)

	_ = json.Unmarshal(dotsByte, &dct)
	src = strings.Split(dots, ",")
	chkRet := false
	if (len(dct) * 2) == len(src) {
		for i, dot := range dct {
			j := i * 2
			k := i*2 + 1
			sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[j]), 64)
			sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[k]), 64)
			// 校验点的位置,在原有的区域上添加额外边距进行扩张计算区域,不推荐设置过大的padding
			// 例如：文本的宽和高为30，校验范围x为10-40，y为15-45，此时扩充5像素后校验范围宽和高为40，则校验范围x为5-45，位置y为10-50
			chkRet = captcha.CheckPointDistWithPadding(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height), 5)
			if !chkRet {
				break
			}
		}
	}
	if chkRet {
		return true
	}
	return false
}
