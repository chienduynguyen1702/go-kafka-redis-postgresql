package main

import (
	"log"
	"os"
	"vcs-kafka-learning-go-gateway/controllers"
	"vcs-kafka-learning-go-gateway/initializers"
)

func init() {
	// load env
	if os.Getenv("RUN_ON") == "localhost" {
		initializers.LoadEnvVariables()
	}
	// Connect to Redis
	redisClient := initializers.ConnectRedis()
	controllers.SetRedis(redisClient)

	// Connect to Kafka and check if working topic is exist
	orderTopicConsumer := initializers.ConnectConsumerToKafka()

	// Create topic Consumer
	controllers.SetKafkaReader(orderTopicConsumer)

	log.Println("Finish initiation !")
}
func main() {
	controllers.StartService()
}
