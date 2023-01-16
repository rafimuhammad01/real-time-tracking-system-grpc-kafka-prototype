package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func initKafka(topic, network, address string, partition int) *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), network, address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return conn
}
