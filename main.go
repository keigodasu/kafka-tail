package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	kafkat "github.com/keigodasu/kafka-tail/lib"
	"os"
)


func main() {
	var bootstrapServers, topic, groupID string

	flag.StringVar(&bootstrapServers, "bootstrap-servers", "", "bootstrap servers")
	flag.StringVar(&topic, "topic", "", "topic (currently doesn't support regular expressions and multiple topics)")
	flag.StringVar(&groupID, "group-id", "", "consumer group id")
	flag.Parse()

	if bootstrapServers == "" {
		fmt.Fprintln(os.Stderr, "'bootstrap-servers' is mandatory")
		os.Exit(1)
	}
	if topic == "" {
		fmt.Fprintln(os.Stderr, "'topic' is mandatory")
		os.Exit(1)
	}
	if groupID == "" {
		u, _ := uuid.NewRandom()
		groupID = u.String()
	}

	client := kafkat.New(bootstrapServers, topic, groupID)
	client.Run()
}
