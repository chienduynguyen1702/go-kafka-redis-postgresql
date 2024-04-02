# API Gateway
## Responsibility
- Receive order requests
- Also push order events to "orders" topic in Kafka as a producer.
- Listen for the "responses" topic from Kafka as a consumer