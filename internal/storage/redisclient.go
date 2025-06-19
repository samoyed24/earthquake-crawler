package storage

import (
	"context"
	"earthquake-crawler/internal/config"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisClientTypeDef struct {
	client *redis.Client
	ctx    context.Context
}

var redisClient = new(RedisClientTypeDef)

func InitRedisClient() error {
	redisClient.ctx = context.Background()
	redisClient.client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Cfg.Redis.Addr, config.Cfg.Redis.Port),
		Password: config.Cfg.Redis.Password,
		DB:       config.Cfg.Redis.DB,
	})

	_, err := redisClient.client.Ping(redisClient.ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func CloseRedisClient() error {
	if redisClient.client == nil {
		return fmt.Errorf("redis client未初始化")
	}
	return redisClient.client.Close()
}

func RPushRedis(key string, value interface{}) error {
	return redisClient.client.RPush(redisClient.ctx, key, value).Err()
}
