list-topics:
	docker-compose exec kafka kafka-topics.sh --list --bootstrap-server kafka:9092
describe-topics:
	docker-compose exec kafka kafka-topics.sh --describe --bootstrap-server kafka:9092
create-orders-topic:
	docker-compose exec kafka kafka-topics.sh --create --topic orders --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092
create-responses-topic:
	docker-compose exec kafka kafka-topics.sh --create --topic responses --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092