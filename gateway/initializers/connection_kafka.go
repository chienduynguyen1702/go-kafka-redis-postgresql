package initializers

import (
	"log"
	"os"
	kafka_connection "vcs-kafka-learning-go-gateway/modules/kafka"

	"github.com/segmentio/kafka-go"
)

const (
	orderTopicName    = "orders"
	responseTopicName = "responses"
)

var kafkaBroker = []string{os.Getenv("KAFKA_BROKER")}

func ConnectKafka() (*kafka.Reader, *kafka.Writer) {
	// Initialize Kafka writer
	kafkaWriter, err := kafka_connection.NewKafkaWriter(orderTopicName, kafkaBroker)
	if err != nil {
		log.Fatalf("Failed to create Kafka writer: %v", err)
	}
	// Initialize Kafka reader
	kafkaReader, err := kafka_connection.NewKafkaReader(responseTopicName, kafkaBroker)
	if err != nil {
		log.Fatalf("Failed to create Kafka reader: %v", err)
	}
	return kafkaReader, kafkaWriter
}
