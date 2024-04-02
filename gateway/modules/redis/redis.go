package redis_connection

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(redisOptions *redis.Options, ctx context.Context) *redis.Client {
	rdb := redis.NewClient(redisOptions)

	val, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v. Redis PING response: %s", err, val)
	}

	return rdb
}
