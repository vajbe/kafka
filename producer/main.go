package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}
	defer p.Close()

	topic := "test_topic"

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Message %d", i)
		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(msg),
		}, nil)

		if err != nil {
			log.Printf("Failed to send message: %s\n", err)
		} else {
			fmt.Printf("Produced: %s\n", msg)
		}

		time.Sleep(time.Second)
	}
}
