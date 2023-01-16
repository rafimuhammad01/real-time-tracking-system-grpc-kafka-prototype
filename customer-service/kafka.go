package main

import (
	"github.com/segmentio/kafka-go"
)

func NewKafkaReader(topic, network string, address []string, partition int) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  address,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return r
}
