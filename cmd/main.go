package main

import (
	"flag"
	kafkat "github.com/keigodasu/kafka-tail"
)

func main() {
	var bootstrapServers, topic string

	flag.StringVar(&bootstrapServers, "bootstrap-servers", "", "bootstrap servers")
	flag.StringVar(&topic, "topic", "", "topic")
	flag.Parse()

	if bootstrapServers == "" {
		panic("bootstrap-servers is needed")
	}
	if topic == "" {
		panic("topic is needed")
	}

	client := kafkat.New(bootstrapServers, topic)
	client.Run()
}