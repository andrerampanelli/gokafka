package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer := NewKafkaProducer()
	defer producer.Close()

	deliveryChan := make(chan kafka.Event)
	go DeliveryReport(deliveryChan)

	// key := []byte("transfer") // This key will be used to determine the partition that the message will be sent to or send nil to let Kafka decide.
	// err := Publish("Hello, Kafka!", "test", producer, key, deliveryChan)

	err := Publish("Hello, Kafka!", "test", producer, nil, deliveryChan)
	if err != nil {
		panic(err)
	}
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "gokafka-kafka-1:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true", // If you want to ensure that the messages are delivered in order and exactly once, set this to true. If true then acks must be all.
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}
	return p
}

func Publish(message string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	defer close(deliveryChan)

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Key:            key,
	}

	err := producer.Produce(msg, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				panic(ev.TopicPartition.Error)
			} else {
				println("Message delivered to topic ", *ev.TopicPartition.Topic, " partition ", ev.TopicPartition.Partition)
			}
		}
	}
}
