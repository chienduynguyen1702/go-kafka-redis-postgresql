BROKER_HOST=broker
KAFKA_PATH_SH=/opt/kafka/bin

list-topics:
	docker-compose exec $(BROKER_HOST) $(KAFKA_PATH_SH)/kafka-topics.sh --list --bootstrap-server $(BROKER_HOST):9092
describe-topics:
	docker-compose exec $(BROKER_HOST) $(KAFKA_PATH_SH)/kafka-topics.sh --describe --bootstrap-server $(BROKER_HOST):9092
create-orders-topic:
	docker-compose exec $(BROKER_HOST) $(KAFKA_PATH_SH)/kafka-topics.sh --create --topic orders --partitions 1 --replication-factor 1 --bootstrap-server $(BROKER_HOST):9092
create-responses-topic:
	docker-compose exec $(BROKER_HOST) $(KAFKA_PATH_SH)/kafka-topics.sh --create --topic responses --partitions 1 --replication-factor 1 --bootstrap-server $(BROKER_HOST):9092
consume-orders-topic:
	docker-compose exec $(BROKER_HOST) $(KAFKA_PATH_SH)/kafka-console-consumer.sh --topic orders --bootstrap-server $(BROKER_HOST):9092 --from-beginning
produce-orders-topic:
	docker-compose exec $(BROKER_HOST) $(KAFKA_PATH_SH)/kafka-console-producer.sh --topic orders --bootstrap-server $(BROKER_HOST):9092
