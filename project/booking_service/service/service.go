package service

import "github.com/nislovskaya/microservice_architecture/project/booking_service/service/booking"

type Service struct {
	booking.Service
}

func New(opts ...Option) *Service {
	service := &Service{}

	for _, option := range opts {
		option(service)
	}

	return service
}
