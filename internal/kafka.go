package internal

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	l *log.Logger
	c *kafka.Consumer
}

func NewKafkaConsumer(l *log.Logger) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	kafkaConsumer := &KafkaConsumer{l, c}

	return kafkaConsumer, nil
}

func (k *KafkaConsumer) SubscribeToTopic(topic string, ch chan *kafka.Message) error {
	err := k.c.Subscribe(topic, nil)

	if err != nil {
		return err
	}

	k.l.Printf("[Kafka Consumer] Subscribed to topic %s\n", topic)

	go func() {
		for {
			msg, _ := k.c.ReadMessage(-1)
			ch <- msg
		}
	}()

	return nil
}
