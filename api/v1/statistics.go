package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"qiublog/db"
	"qiublog/model"
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
		err          error
		articleCount int64   //文章总数
		mainUV       int64   //浏览量
		wordsTotal   float64 //文章总字数
		elapsedTime  int64   //建站时间
		lastUpdated  int64   //最后更新时间
	)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	articleCount, err = db.Rdb.Get(ctx, "main:ac").Int64()
	if err != nil {
		model.Db.Model(model.Article{}).Count(&articleCount)
		db.Rdb.Set(ctx, "main:ac", articleCount, 3*24*time.Hour)
	}
	mainUV, _ = db.Rdb.PFCount(ctx, "main:uv").Result()
	wordsTotal, err = db.Rdb.Get(ctx, "main:wt").Float64()
	if err != nil {
		sql := fmt.Sprintf("select ROUND(DATA_LENGTH/80000,2) from information_schema.TABLES where table_schema='%s' and table_name='article';", utils.Config.Database.DbName)
		model.Db.Raw(sql).Scan(&wordsTotal)
		db.Rdb.Set(ctx, "main:wt", wordsTotal, 3*24*time.Hour)
	}
	lastUpdated, err = db.Rdb.Get(ctx, "main:lut").Int64()
	if err != nil {
		var ar model.Article
		model.Db.Model(model.Article{}).Order("updated_at desc").Take(&ar)
		lastUpdated = ar.UpdatedAt.Unix()
		db.Rdb.Set(ctx, "main:lut", lastUpdated, 3*24*time.Hour)
	}
	elapsedTime = utils.Config.SiteInfo.ConstructionTime
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
