package routing

import (
	"github.com/nislovskaya/microservice_architecture/project/route_service/repository"
	"github.com/sirupsen/logrus"
)

type Option func(r *routing)

func WithLogger(logger *logrus.Entry) Option {
	return func(r *routing) {
		r.logger = logger
	}
}

func WithRepository(repo repository.Repository) Option {
	return func(r *routing) {
		r.repo = repo
	}
}
