package initializers

import (
	"context"
	"log"
	"os"
	kafka_connection "vcs-kafka-learning-go-gateway/modules/kafka"

	"github.com/segmentio/kafka-go"
)

const (
	orderTopicName = "orders"
)

var kafkaBroker = []string{os.Getenv("KAFKA_BROKER")}
var kafkaAddress = os.Getenv("KAFKA_BROKER")
var orderTopicConsumer *kafka.Reader

// Connect to Kafka and check if working topic is exist
func ConnectConsumerToKafka() *kafka.Reader {
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaAddress, orderTopicName, 0)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	if err := CreateOrderTopicIfNotExists(conn); err != nil {
		panic(err.Error())
	}
	orderTopicConsumer = kafka_connection.NewKafkaReader(orderTopicName, kafkaBroker)
	return orderTopicConsumer
}

// CreateTopicIfNotExists create topic if not exists
func CreateOrderTopicIfNotExists(conn *kafka.Conn) error {
	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	for _, p := range partitions {
		if p.Topic == orderTopicName {
			log.Printf("Topic %s is already existed.", orderTopicName)
			return nil
		}
	}

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             orderTopicName,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Topic %s created.", orderTopicName)
	return nil
}
