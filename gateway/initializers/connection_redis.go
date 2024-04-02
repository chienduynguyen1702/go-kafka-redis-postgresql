package initializers

import (
	"context"
	"fmt"
	"os"
	"strconv"

	redis_connection "vcs-kafka-learning-go-gateway/modules/redis"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	redisAddr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisProtocol, _ := strconv.Atoi(os.Getenv("REDIS_PROTOCOL"))

	redisOptions := &redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
		Protocol: redisProtocol,
	}
	return redis_connection.NewRedisClient(redisOptions, context.Background())
}
