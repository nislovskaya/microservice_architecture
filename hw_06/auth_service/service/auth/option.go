package auth

import (
	"github.com/go-redis/redis/v8"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/kafka"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/repository"
	"github.com/sirupsen/logrus"
)

type Option func(fs *auth)

func WithLogger(logger *logrus.Entry) Option {
	return func(a *auth) {
		a.logger = logger
	}
}

func WithRepository(repo repository.Repository) Option {
	return func(a *auth) {
		a.repo = repo
	}
}

func WithSecretKey(secretKey string) Option {
	return func(a *auth) {
		a.secretKey = secretKey
	}
}

func WithKafkaProducer(producer *kafka.Producer) Option {
	return func(a *auth) {
		a.kafka = producer
	}
}

func WithRedis(redis *redis.Client) Option {
	return func(a *auth) {
		a.redis = redis
	}
}
