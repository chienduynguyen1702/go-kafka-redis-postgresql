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
	RedisClient            *redis.Client
	ctxBackground          = context.Background()
)

func SetRedis(rc *redis.Client) {
	RedisClient = rc
	log.Println("Connected to redis !")
}
func SetKafkaWriter(kw *kafka.Writer) {
	OrdersTopicProducer = kw
	log.Println("Connected to kafka as Producer !")
	// print the kafka writer
	// log.Printf("Kafka Writer: %v", OrdersTopicProducer)
}
func SetKafkaReader(kr *kafka.Reader) {
	ResponsesTopicConsumer = kr
	log.Println("Connected to kafka as Consumer !")
	// print the kafka reader
	// log.Printf("Kafka Reader: %v", ResponsesTopicConsumer)
}
