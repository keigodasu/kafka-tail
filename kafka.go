package kafka_tail

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Client struct {
	kafka *kafka.Consumer
	topic string
}

func New(bootstrapServer string, topic string) *Client {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServer,
		"group.id": "testGroup",
	})

	if err != nil {
		panic(err)
	}

	return &Client{
		kafka: c,
		topic: topic,
	}
}

func (c *Client) Run()  {
	c.kafka.SubscribeTopics([]string{c.topic}, nil)
	fmt.Println("start consuming...")
	for  {
		msg, err := c.kafka.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}