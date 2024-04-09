package main

import (
	"log"
	"os"
	"vcs-kafka-learning-go-gateway/controllers"
	"vcs-kafka-learning-go-gateway/docs"
	"vcs-kafka-learning-go-gateway/initializers"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
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
	orderTopicProducer := initializers.ConnectProducerToKafka()

	// Create topic producer
	controllers.SetKafkaWriter(orderTopicProducer)

	log.Println("Finish initiation !")
}
func main() {

	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/orders", controllers.CreateOrder)
		api.GET("/items", controllers.GetItem)
		api.POST("/items", controllers.ResetItemQuantity)

	}

	docs.SwaggerInfo.Title = "kafka Backend API"
	docs.SwaggerInfo.Description = "This is a API documentation for kafka Backend."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	// docs.SwaggerInfo.BasePath = "/api/v1"
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.Printf("Server is running on :8080 with %s Gin mode", os.Getenv("GIN_MODE"))

	go controllers.SetRPCClient()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
