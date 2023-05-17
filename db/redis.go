package db

import (
	"context"
	"fmt"
	redis "github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"qiublog/utils"
	"time"
)

var Rdb *redis.Client

func InitRedis() {
	log.Debug("init redis...")
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.Config.Server.Database.Redis.Host, utils.Config.Server.Database.Redis.Port),
		Password: utils.Config.Server.Database.Redis.Password, // no password set
		DB:       utils.Config.Server.Database.Redis.Db,       // use default DB
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := Rdb.Ping(ctx).Result()
	//fmt.Println(pong, err)
	if err != nil {
		panic(fmt.Sprintf("连接redis出错，错误信息：%v", err))
	}
	// 开发环境，清空缓存
	//if utils.IsDev() {
	//	Rdb.FlushAll(ctx)
	//}
	log.Debug("init redis success!")
}

// Allow 通过redis的value判断第几次访问并返回是否允许访问
func Allow(key string, events int64, per time.Duration) bool {
	ctx := context.Background()
	curr := Rdb.LLen(ctx, key).Val()
	if curr >= events {
		return false
	}
	if v := Rdb.Exists(ctx, key).Val(); v == 0 {
		pipe := Rdb.TxPipeline()
		pipe.RPush(ctx, key, key)
		//设置过期时间
		pipe.Expire(ctx, key, per)
		_, _ = pipe.Exec(ctx)
	} else {
		Rdb.RPushX(ctx, key, key)
	}
	return true
}
