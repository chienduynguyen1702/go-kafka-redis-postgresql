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

	// Write the message to the Kafka topic
	err = KafkaWriter.WriteMessages(c, message)
	if err != nil {
		log.Printf("Error writing message to Kafka: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Read the response from the Kafka topic
	m, err := KafkaReader.ReadMessage(c)
	if err != nil {
		log.Printf("Error reading message from Kafka: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Process the response message
	var response models.Response
	err = json.Unmarshal(m.Value, &response)
	if err != nil {
		log.Printf("Error unmarshalling order: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Return the status of the order to the client
	c.JSON(http.StatusOK, gin.H{"status": response})
}