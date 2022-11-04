package repository

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaConfig = &kafka.ConfigMap{
	"bootstrap.servers": "localhost:9092",
	"group.id":          "myGroup",
	"auto.offset.reset": "earliest",
}

// Kafka struct
type Kafka struct {
	Producer *kafka.Producer
}

// NewKafka func
func NewKafka() *Kafka {
	return &Kafka{}
}

// Consume topic
func (k *Kafka) Consume(topic string, fromBegining bool) {
	c, err := kafka.NewConsumer(kafkaConfig)

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			println("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			println("Consumer error: %v (%v)\n", err, msg)
		}
	}

}
