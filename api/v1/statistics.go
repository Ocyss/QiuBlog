package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"qiublog/db"
	"qiublog/utils"
	"qiublog/utils/errmsg"
	"time"
)

/*
	访问量：每次访问 +1
	浏览量：文章浏览60% +1
*/

func MainSetUV(c *gin.Context) (int, any) {
	ctx := context.Background()
	err := db.Rdb.PFAdd(ctx, "main:uv", c.ClientIP()).Err()
	if err != nil {
		return errmsg.ERROR, nil
	}
	return errmsg.SUCCESS, nil
}

func GetStatistics(c *gin.Context) (int, any) {
	var (
		articleCount int64
		mainUV       int64
		wordsTotal   int64
		elapsedTime  int64
		lastUpdated  int64
	)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	articleCount, _ = db.Rdb.Get(ctx, "main:ac").Int64()
	mainUV, _ = db.Rdb.PFCount(ctx, "main:uv").Result()
	wordsTotal, _ = db.Rdb.Get(ctx, "main:wt").Int64()
	lastUpdated, _ = db.Rdb.Get(ctx, "main:lut").Int64()
	elapsedTime = utils.Config.ConstructionTime
	return errmsg.SUCCESS, gin.H{
		"data": gin.H{
			"article_count": articleCount,
			"main_uv":       mainUV,
			"words_total":   wordsTotal,
			"elapsed_time":  elapsedTime,
			"last_updated":  lastUpdated,
		},
	}
}
