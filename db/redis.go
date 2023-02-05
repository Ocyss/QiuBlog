package db

import (
	"context"
	"fmt"
	redis "github.com/redis/go-redis/v9"
	"qiublog/utils"
)

var Rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.Config.Redis.RedisHost, utils.Config.Redis.RedisPort),
		Password: utils.Config.Redis.RedisPassword, // no password set
		DB:       utils.Config.Redis.RedisDb,       // use default DB
	})

	_, err := Rdb.Ping(ctx).Result()
	//fmt.Println(pong, err)
	if err != nil {
		panic(fmt.Sprintf("连接redis出错，错误信息：%v", err))
	}
}
