package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"vcs-kafka-learning-go-gateway/models"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

// CreateOrder godoc
//
// @Summary Create an order
// @Schemes
// @Description  Create an order
// @Tags Orders
// @Accept json
// @Produce json
// @Param body body models.CreateOrderRequest true "The request body."
// @Success 200 {string}  string "Order have been created successfully."
// @Router /api/orders [post]
func CreateOrder(c *gin.Context) {
	// Parse order data from request
	var orderData models.Order
	if err := c.BindJSON(&orderData); err != nil {
		log.Panic(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Marshal the order data into JSON
	orderJSON, err := json.Marshal(orderData)
	if err != nil {
		log.Printf("Error marshalling order: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Create a Kafka message with the order data
	message := kafka.Message{
		Key:   nil,       // Optionally, you can specify a message key
		Value: orderJSON, // Set the value of the message to the marshalled order JSON
	}
	// startTime := time.Now()

	// Write the message to the Kafka topic
	err = OrdersTopicProducer.WriteMessages(c, message)
	if err != nil {
		log.Printf("Error writing message to Kafka: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	log.Println("Send to broker message: ", orderData)

	// Read the response from responseChannel
	response := <-responseChannel
	log.Println("Response from rpc client: ", response)
	if response.IsSucceed {
		c.JSON(http.StatusOK, gin.H{"status": "Order created successfully"})
		return
	}
	// If the response is not successful, return an error to the client
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
}
