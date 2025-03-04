package service

import "github.com/nislovskaya/microservice_architecture/project/route_service/service/routing"

type Service struct {
	routing.Service
}

func New(opts ...Option) *Service {
	service := &Service{}

	for _, option := range opts {
		option(service)
	}

	return service
}
