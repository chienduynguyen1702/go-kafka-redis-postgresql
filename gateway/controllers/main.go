package controllers

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	OrdersTopicProducer    *kafka.Writer
	ResponsesTopicConsumer *kafka.Reader
	KafkaConnection        *kafka.Conn
	RedisClient            *redis.Client
	ctxBackground          = context.Background()
)

func SetRedis(rc *redis.Client) {
	RedisClient = rc
	log.Println("Connected to redis !")
}
func SetKafkaConnection(kc *kafka.Conn) {
	KafkaConnection = kc
	log.Println("Connected to kafka !")
}

func SetKafkaReader(kr *kafka.Reader) {
	ResponsesTopicConsumer = kr
	log.Println("Connected to kafka as Consumer !")
}

func SetKafkaWriter(kw *kafka.Writer) {
	OrdersTopicProducer = kw
	log.Println("Connected to kafka as Producer !")
}
