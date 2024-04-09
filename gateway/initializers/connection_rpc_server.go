package initializers

import (
	"log"
	"vcs-kafka-learning-go-gateway/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateRpcServer(host string, port string) proto.ResponseServiceClient {
	endpoint := host + ":" + port
	log.Print("Connecting to gRPC server at ", endpoint)
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewResponseServiceClient(conn)
	return client
}
