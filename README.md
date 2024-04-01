# Sale Campaigne Simulation
## Tech stack
- Golang
- Kafka
- Redis
- PostgreSQL
- Docker

## Architecture

The architecture is composed of 4 main components:
- client : make request to sale serivce by gRPC
- server : running sale service server, push data into kafka and redis.
- Kafka : make  sure the message will be delivered in order. It's used for communication between client and server.
- Redis :  store the real time stock information for each product.
- PostgreSQL : store historical stock information for each product.

## Start:

```bash
docker-compose up --build -d
```

