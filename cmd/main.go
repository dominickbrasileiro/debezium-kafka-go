package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dominickbrasileiro/debezium-kafka-go/internal"
)

func init() {
	err := internal.CheckDebeziumConnector()

	if err != nil {
		panic(err)
	}
}

func main() {
	topics := []string{
		"postgres.public.product",
		"postgres.public.student",
	}

	logger := log.Default()

	for _, topic := range topics {
		kc, err := internal.NewKafkaConsumer(logger)

		if err != nil {
			panic(err)
		}

		msgChannel := make(chan *kafka.Message)

		kc.SubscribeToTopic(topic, msgChannel)

		go func() {
			for msg := range msgChannel {
				logger.Println("--------------------------------------")
				logger.Println("New message:", msg)
				logger.Println("Data:", string(msg.Value))
			}
		}()
	}

	// Handle exit signal (CTRL + C)
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, os.Interrupt)
	<-exitChannel

	logger.Println()
}
