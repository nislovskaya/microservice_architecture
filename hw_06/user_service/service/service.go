package service

import "github.com/nislovskaya/microservice_architecture/hw_04/crud_app/service/user"

type Service struct {
	user.Service
}

func New(opts ...Option) *Service {
	service := &Service{}

	for _, option := range opts {
		option(service)
	}

	return service
}
