package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(bootstrapServers string) (*Producer, error) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(bootstrapServers),
		Balancer: &kafka.LeastBytes{},
	}

	return &Producer{writer: w}, nil
}

func (p *Producer) Publish(topic string, message interface{}) error {
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	err = p.writer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Value: msgBytes,
	})
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
