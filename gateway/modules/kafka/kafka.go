package kafka_connection

import (
	"github.com/segmentio/kafka-go"
)

// const (
// 	kafkaBroker        = "localhost:9092"
// 	kafkaOrderTopic    = "orders"
// 	kafkaResponseTopic = "responses"
// )

// Write message to Kafka topic
func NewKafkaWriter(topic string, kafkaBroker []string) (*kafka.Writer, error) {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: kafkaBroker,
		Topic:   topic,
	}), nil
}

func NewKafkaReader(topic string, kafkaBroker []string) (*kafka.Reader, error) {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   kafkaBroker,
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	}), nil
}
