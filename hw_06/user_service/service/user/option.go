package user

import (
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/kafka"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/repository"
	"github.com/sirupsen/logrus"
)

type Option func(fs *user)

func WithLogger(logger *logrus.Entry) Option {
	return func(u *user) {
		u.logger = logger
	}
}

func WithRepository(repo repository.Repository) Option {
	return func(u *user) {
		u.repo = repo
	}
}

func WithKafkaConsumer(consumer *kafka.Consumer) Option {
	return func(u *user) {
		u.consumer = consumer
	}
}
