package kafka_connection

import (
	"github.com/segmentio/kafka-go"
)

// const (
// 	kafkaAddress        = "localhost:9092"
// 	kafkaOrderTopic    = "orders"
// 	kafkaResponseTopic = "responses"
// )

// Write message to Kafka topic
func NewKafkaWriter(topic string, kafkaAddress string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// Read message from Kafka topic
func NewKafkaReader(topic string, kafkaBroker []string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  kafkaBroker,
		Topic:    topic,
		GroupID:  "vcs-kafka-learning-go-gateway",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}
