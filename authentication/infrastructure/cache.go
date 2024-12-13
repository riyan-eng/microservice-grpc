package infrastructure

import (
	"context"
	"fmt"
	"log"

	"server/env"

	"github.com/redis/go-redis/v9"
)

// var Redis *redis.Client

func ConnRedis() (*redis.Client, error) {
	addr := fmt.Sprintf("%v:%v", env.NewEnv().REDIS_HOST, env.NewEnv().REDIS_PORT)
	cache := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: env.NewEnv().REDIS_USERNAME,
		Password: env.NewEnv().REDIS_PASSWORD,
		DB:       env.NewEnv().REDIS_DATABASE,
	})
	ctx := context.Background()
	if err := cache.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis: can't ping to redis - %v", err)
	}
	log.Println("redis: connection opened")
	return cache, nil
}
