package repository

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaConfig = &kafka.ConfigMap{
	"bootstrap.servers": "localhost:9092",
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
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition.Offset, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

}

// Produce topic
func Produce(topic string, message string) {
	p, err := kafka.NewProducer(kafkaConfig)

	if err != nil {
		panic(err)
	}

	defer p.Close()

	p.ProduceChannel() <- &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}

	e := <-p.Events()

	switch ev := e.(type) {
	case *kafka.Message:
		if ev.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
		}

	}
}
