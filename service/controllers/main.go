package controllers

import (
	"context"
	"log"
	"vcs-kafka-learning-go-gateway/proto"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	OrdersTopicConsumer *kafka.Reader
	KafkaConnection     *kafka.Conn
	RedisClient         *redis.Client
	// ctxBackground       = context.Background()
	rpcClient           proto.ResponseServiceClient
	ctxBackground, cancel = context.WithCancel(context.Background())
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
	OrdersTopicConsumer = kr
	log.Println("Connected to kafka as Consumer !")
}

func SetRpcClient(c proto.ResponseServiceClient) {
	rpcClient = c
	log.Println("Connected to RPC server !")
}
