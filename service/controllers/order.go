package controllers

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func StartService() {
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

		// Read message from Kafka
		msg, err := OrdersTopicConsumer.ReadMessage(ctxBackground)
		if err != nil {
			log.Printf("Error reading message from Kafka: %v", err)
			continue
		}

		// Process the received message (e.g., print to console)
		fmt.Printf("Received message: %s\n", msg.Value)
	}
}
