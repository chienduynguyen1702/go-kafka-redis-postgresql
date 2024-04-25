package controllers

import (
	"context"
	"log"
	"net"
	"os"
	"vcs-kafka-learning-go-gateway/models"
	"vcs-kafka-learning-go-gateway/proto"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

type responseServer struct {
	proto.ResponseServiceServer
}

var (
	// ReponseRpcServerHost = os.Getenv("RPC_HOST")
	ReponseRpcServerPort = ":" + os.Getenv("RPC_PORT")

	OrdersTopicProducer    *kafka.Writer
	ResponsesTopicConsumer *kafka.Reader
	KafkaConnection        *kafka.Conn
	RedisClient            *redis.Client
	ctxBackground          = context.Background()
	responseChannel        = make(chan models.Response)
	ResponseServer         responseServer
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
func SetRPCClient() {
	lis, err := net.Listen("tcp", ReponseRpcServerPort)
	if err != nil {
		panic(err)
	}
	rpcServer := grpc.NewServer()
	proto.RegisterResponseServiceServer(rpcServer, &ResponseServer)
	log.Printf("Response rpc server listening at %s", lis.Addr())
	if err := rpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
