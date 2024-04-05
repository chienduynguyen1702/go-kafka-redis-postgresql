package kafka_connection

import (
	"github.com/segmentio/kafka-go"
)

func NewConn(kafkaAddress string) *kafka.Conn {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	return conn
}

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
