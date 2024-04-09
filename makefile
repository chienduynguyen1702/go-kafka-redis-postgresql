BROKER_HOST=broker
KAFKA_PATH_SH=/opt/kafka/bin
PROJECT_NAME=vcs2
SERVICE=service
GATEWAY=gateway
BROKER=broker
REDIS=redis

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
create-pb:
	protoc --go_out=. --go-grpc_out=. proto/response.proto
restart-compose:
	docker-compose down
	docker-compose up --build -d
reset-compose:
	docker-compose down
	docker volume prune -f
	docker-compose up --build -d
log-compose:
	docker-compose logs -f --tail 30 -t
log-service:
	docker-compose logs -f --tail 30 $(SERVICE)
log-gateway:
	docker-compose logs -f --tail 30 $(GATEWAY)
log-broker:
	docker-compose logs -f --tail 30 $(BROKER)	


delete-order-topic:
	docker-compose exec $(BROKER_HOST) $(KAFKA_PATH_SH)/kafka-topics.sh --delete --topic orders --bootstrap-server $(BROKER_HOST):9092
delete-__consumer_offsets-topic:
	docker-compose exec $(BROKER_HOST) $(KAFKA_PATH_SH)/kafka-topics.sh --delete --topic __consumer_offsets --bootstrap-server $(BROKER_HOST):9092



start-compose:
	docker compose up --build -d
