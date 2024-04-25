package initializers

import (
	"log"
	"os"
	"vcs-kafka-learning-go-gateway/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ReponseRpcServerHost = os.Getenv("RPC_HOST")
	ReponseRpcServerPort = os.Getenv("RPC_PORT")
)

func ConnectToRPCServer() proto.ResponseServiceClient {
	endpoint := ReponseRpcServerHost + ":" + ReponseRpcServerPort
	log.Print("Connecting to gRPC server at ", endpoint)
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewResponseServiceClient(conn)
	return client
}
