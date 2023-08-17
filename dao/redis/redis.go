package redis

import (
	"app/settings"
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

// Init 初始化redis连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("connect redis failed", zap.Error(err))
		return
	}
	return
}
