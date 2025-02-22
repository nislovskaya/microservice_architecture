package service

import (
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/service/auth"
)

type Service struct {
	auth.Service
}

func New(opts ...Option) *Service {
	service := &Service{}

	for _, option := range opts {
		option(service)
	}

	return service
}
