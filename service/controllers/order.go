package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vcs-kafka-learning-go-gateway/models"
	"vcs-kafka-learning-go-gateway/proto"
)

func StartService() {
	log.Printf("Starting the consumer...")
	// Infinite loop to keep the consumer alivesigChan := make(chan os.Signal, 1)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Received termination signal. Stopping the consumer...")
		cancel()
	}()
	for {
		// Check if context is cancelled
		select {
		case <-ctxBackground.Done():
			log.Println("Consumer stopped.")
			return
		default:
		}
		log.Printf("in for consumer...")

		// Read message from Kafka
		msg, err := OrdersTopicConsumer.ReadMessage(ctxBackground)
		if err != nil {
			log.Printf("Error reading message from Kafka: %v", err)
			continue
		}

		// Process the received message (e.g., print to console)
		fmt.Printf("Received message: %s\n", msg.Value)
		order := models.Order{}
		if err := json.Unmarshal(msg.Value, &order); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		// check
		fmt.Printf("Order ID: %d, Name: %s, Quantity: %d\n", order.ID, order.Name, order.Quantity)

		// Process the received message (e.g., save to Redis)
		// Get currentQuantity of food from Redis
		key := fmt.Sprintf("food:%s:quantity", order.Name)
		currentQuantity, err := RedisClient.HGet(ctxBackground, key, "quantity").Int()
		if err != nil {
			log.Println("Error getting quantity from Redis: ", err)
			_, err = (rpcClient).SendResponse(ctxBackground, &proto.ResponseBody{
				Message:   "Internal Server Error",
				IsSucceed: false,
			})
			if err != nil {
				log.Println("Error sending response to RPC server: ", err)
			}
			continue
		}

		// Check if current quantity is less than order quantity
		if currentQuantity < order.Quantity {
			_, err = (rpcClient).SendResponse(ctxBackground, &proto.ResponseBody{
				Message:   "Invalid quantity:invalid quantity",
				IsSucceed: false,
			})
			if err != nil {
				log.Println("Error sending response to RPC server: ", err)
			}
			continue
		}

		// Check if current quantity is 0
		if currentQuantity == 0 {
			_, err = (rpcClient).SendResponse(ctxBackground, &proto.ResponseBody{
				Message:   "Out of stock",
				IsSucceed: false,
			})
			if err != nil {
				log.Println("Error sending response to RPC server: ", err)
			}
			continue
		}

		// Calculate new quantity after subtracting order quantity
		newQuantity := currentQuantity - order.Quantity

		// Update quantity in Redis
		err = RedisClient.HSet(ctxBackground, key, "quantity", newQuantity).Err()
		if err != nil {
			// Handle error
			log.Printf("Error updating quantity for %s: %v", order.Name, err)
			_, err = (rpcClient).SendResponse(ctxBackground, &proto.ResponseBody{
				Message:   "Internal Server Error",
				IsSucceed: false,
			})
			if err != nil {
				log.Println("Error sending response to RPC server: ", err)
			}
			continue
		}

		// Process the received message (e.g., send to RPC server)
		_, err = (rpcClient).SendResponse(ctxBackground, &proto.ResponseBody{
			Message:   "Order placed successfully",
			IsSucceed: true,
		})
		if err != nil {
			log.Println("Error sending response to RPC server: ", err)
		}
	}
}
