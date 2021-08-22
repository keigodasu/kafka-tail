package kafka_tail

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Client struct {
	kafka *kafka.Consumer
	topic string
}

type Record struct {
	messageTimestamp string
	topic string
	topicPartition int32
	offset string
	key string
	message string
}

func New(bootstrapServer string, topic string, groupID string) *Client {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServer,
		"group.id": groupID,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start with: %s", err)
		os.Exit(1)
	}

	return &Client{
		kafka: c,
		topic: topic,
	}
}

func (c *Client) Run()  {
	c.kafka.SubscribeTopics([]string{c.topic}, nil)
	fmt.Println("start consuming...")

	go func() {
		trap := make(chan os.Signal, 1)
		signal.Notify(trap, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
		_ = <-trap
		c.kafka.Close()
		os.Exit(0)
	}()

	for  {
		msg, err := c.kafka.ReadMessage(-1)
		topic := *msg.TopicPartition.Topic
		if err == nil {
			record := Record{
				messageTimestamp: msg.Timestamp.Format(time.RFC3339),
				topic: topic,
				topicPartition:   msg.TopicPartition.Partition,
				offset:           msg.TopicPartition.Offset.String(),
				key:              string(msg.Key),
				message:          string(msg.Value),
			}
			fmt.Printf("%+v\n", record)
		} else {
			fmt.Printf("got error: %v (%v)\n", err, msg)
		}
	}
}