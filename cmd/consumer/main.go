package main

import "github.com/confluentinc/confluent-kafka-go/kafka"

func main() {
	consumer := NewKafkaConsumer([]string{"test"})
	defer consumer.Close()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			println("Message on ", msg.TopicPartition.String(), " at offset (", msg.TopicPartition.Offset, ") with key ()", string(msg.Key), ") and value ()", string(msg.Value), ")")
		} else {
			println("Consumer error: ", err)
		}
	}
}

func NewKafkaConsumer(topics []string) *kafka.Consumer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "gokafka-kafka-1:9092",
		"group.id":          "gokafka-group",
		"auto.offset.reset": "earliest", // If there is no initial offset in Kafka or if the current offset does not exist any more on the server, the consumer will start reading from the latest records.
	}

	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		panic(err)
	}
	return c
}
