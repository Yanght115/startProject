package gredis

import (
	"plan/pkg/config"
	"plan/pkg/logging"

	"github.com/go-redis/redis"
)

var (
	// RedisClient to operate redis
	RedisClient *redis.Client
)

// InitRedisClient init redis client
func InitRedisClient() {
	config := config.GlobalConfig.Redis
	option := redis.Options{
		Addr:       config.Addr,
		Password:   config.Password,
		DB:         config.DB,
		MaxRetries: config.MaxRetries,
	}
	RedisClient = redis.NewClient(&option)
	err := RedisClient.Ping().Err()
	if err != nil {
		logging.Logger.Criticalf("Init Redis Client Error: %v", err)
		panic(err.Error())
	}
}
