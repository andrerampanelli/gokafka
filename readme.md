# Gokafka

Gokafka is a simple Kafka producer and consumer written in Go.

## How to run

### Docker

```bash
docker-compose up
```

### Producer

```bash
go run cmd/producer/main.go
```

### Consumer

```bash
go run cmd/consumer/main.go
```

## Kafka

### Create topic

```bash
docker-compose exec gokafka-kafka-1 kafka-topics --create --bootstrap-server localhost:9092 --partitions 3 --topic test
```

### List topics

```bash
docker-compose exec gokafka-kafka-1 kafka-topics --list --bootstrap-server localhost:9092
```

### Describe topic

```bash
docker-compose exec gokafka-kafka-1 kafka-topics --describe --bootstrap-server localhost:9092 --topic test
```

### Consume messages

```bash
docker-compose exec gokafka-kafka-1 kafka-console-consumer --bootstrap-server localhost:9092 --topic test
```

### Describe consumer group

```bash
docker-compose exec gokafka-kafka-1 kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group gokafka-group
```
