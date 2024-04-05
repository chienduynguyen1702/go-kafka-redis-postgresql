package controllers

import "log"

func StartService() {
	for {
		message, err := OrdersTopicConsumer.ReadMessage(ctxBackground)
		if err != nil {
			log.Printf("Error while reading message from Kafka: %v", err)
			continue
		}
		log.Printf("Message received: %v", string(message.Value))
	}
}
