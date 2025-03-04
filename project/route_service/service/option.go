package service

import "github.com/nislovskaya/microservice_architecture/project/route_service/service/routing"

type Option func(fs *Service)

func WithRoutingService(routingService routing.Service) Option {
	return func(s *Service) {
		s.Service = routingService
	}
}
