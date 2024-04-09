package kafka_connection

import (
	"os"
	"strconv"
	"time"

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
	writeTimeoutEnv := os.Getenv("KAFKA_BATCH_TIMEOUT")
	writeTimeout, err := strconv.ParseInt(writeTimeoutEnv, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	writeBatchTimeout := time.Duration(writeTimeout) * time.Nanosecond
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{kafkaAddress},
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: writeBatchTimeout,
	})
}

// Read message from Kafka topic
func NewKafkaReader(topic string, kafkaBroker []string) *kafka.Reader { // Parse the environment variable for batch timeout duration
	batchTimeoutEnv := os.Getenv("KAFKA_BATCH_TIMEOUT")
	batchTimeout, err := strconv.Atoi(batchTimeoutEnv)
	if err != nil {
		panic(err.Error())
	}
	readBatchTimeout := time.Duration(batchTimeout) * time.Nanosecond

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:          kafkaBroker,
		Topic:            topic,
		GroupID:          "vcs-kafka-learning-go-service",
		MinBytes:         10e3, // 10KB
		MaxBytes:         10e6, // 10MB
		ReadBatchTimeout: readBatchTimeout,
		MaxWait:          readBatchTimeout,
	})
}
