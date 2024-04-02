package controllers

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	KafkaWriter   *kafka.Writer
	KafkaReader   *kafka.Reader
	RedisClient   *redis.Client
	ctxBackground = context.Background()
)

func SetRedis(rc *redis.Client) {
	RedisClient = rc
}
func SetKafkaWriter(kw *kafka.Writer) {
	KafkaWriter = kw
}
func SetKafkaReader(kr *kafka.Reader) {
	KafkaReader = kr
}
