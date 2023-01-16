package main

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type producer struct {
	conn *kafka.Conn
}

func (p *producer) SendDriverLocation(location Location) error {
	p.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	data, err := location.toJSON()
	if err != nil {
		return err
	}

	_, err = p.conn.WriteMessages(
		kafka.Message{Value: data},
	)
	if err != nil {
		return err
	}

	return nil
}
