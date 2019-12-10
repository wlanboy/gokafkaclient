package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	var counter uint64

	kafkatopic := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"192.168.178.120:9092"},
		Topic:    "topic-go",
		Balancer: &kafka.LeastBytes{},
	})

	fmt.Println("Producer started")

	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("New message at", t)
				publishMessage(kafkatopic, counter)
			}
		}
	}()

	time.Sleep(1 * time.Minute)
	ticker.Stop()
	done <- true
	fmt.Println("Producer stopped")

	kafkatopic.Close()
}

func publishMessage(kafkatopic *kafka.Writer, counter uint64) {
	message := kafka.Message{
		Key:   []byte(string(counter)),
		Value: []byte("Hello Kafka"),
	}

	kafkatopic.WriteMessages(context.Background(), message)
}
