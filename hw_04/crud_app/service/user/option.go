package user

import (
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/repository"
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
