package initializers

import (
	"fmt"
	"os"
	kafka_connection "vcs-kafka-learning-go-gateway/modules/kafka"

	"github.com/segmentio/kafka-go"
)

const (
	orderTopicName    = "orders"
	responseTopicName = "responses"
)

var kafkaBroker = []string{os.Getenv("KAFKA_BROKER")}
var kafkaAddress = os.Getenv("KAFKA_BROKER")

func ConnectKafka() (*kafka.Reader, *kafka.Writer) {
	// debug
	fmt.Println("Kafka Broker: ", kafkaBroker)
	fmt.Println("Kafka Address: ", kafkaAddress)
	// Initialize Kafka writer
	kafkaWriter := kafka_connection.NewKafkaWriter(orderTopicName, kafkaAddress)
	// Initialize Kafka reader
	kafkaReader := kafka_connection.NewKafkaReader(responseTopicName, kafkaBroker)
	return kafkaReader, kafkaWriter
}
