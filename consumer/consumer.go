package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {

	kafkatopic := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"192.168.178.120:9092"},
		GroupID:  "consumer-group-go",
		Topic:    "topic-go",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	fmt.Println("Consumer started")

	ctx := context.Background()
	for {
		message, err := kafkatopic.FetchMessage(ctx)
		if err != nil {
			break
			fmt.Println("Consumer stopped")
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		kafkatopic.CommitMessages(ctx, message)
	}

	kafkatopic.Close()
	fmt.Println("Consumer closed")
}
