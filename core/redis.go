package core

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"time"
)

// InitConnectRedis 初始化Redis连接
func InitConnectRedis() *redis.Client {
	return ConnectRedisDB(1)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	// 新建redis连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		DB:       db,
		Password: redisConf.Password,
		PoolSize: redisConf.PoolSize,
	})
	// 设置连接redis时的超时时间
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	// 超时释放资源
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Errorf("redis connect err:%v host:%d", err, redisConf.Addr())
		return nil
	}
	return rdb
}
