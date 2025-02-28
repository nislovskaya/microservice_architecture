package service

import (
	"github.com/nislovskaya/microservice_architecture/hw_06/jwt_verifier/service/validation"
)

type Service struct {
	validation.Service
}

func New(opts ...Option) *Service {
	service := &Service{}

	for _, option := range opts {
		option(service)
	}

	return service
}
