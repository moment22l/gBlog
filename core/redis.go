package core

import (
	"context"
	"gBlog/global"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

// ConnectRedis 连接redis
func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Conf.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Errorf("redis连接失败 %s", redisConf.Addr())
		return nil
	}
	return rdb
}
