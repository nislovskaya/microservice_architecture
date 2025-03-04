package booking

import (
	"github.com/nislovskaya/microservice_architecture/project/booking_service/kafka"
	"github.com/nislovskaya/microservice_architecture/project/booking_service/repository"
	"github.com/sirupsen/logrus"
)

type Option func(b *booking)

func WithLogger(logger *logrus.Entry) Option {
	return func(b *booking) {
		b.logger = logger
	}
}

func WithRepository(repo repository.Repository) Option {
	return func(b *booking) {
		b.repo = repo
	}
}

func WithKafkaConsumer(consumer *kafka.Consumer) Option {
	return func(b *booking) {
		b.consumer = consumer
	}
}
