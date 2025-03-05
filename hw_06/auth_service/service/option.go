package service

import "github.com/nislovskaya/microservice_architecture/hw_06/auth_service/service/auth"

type Option func(fs *Service)

func WithAuthService(authService auth.Service) Option {
	return func(s *Service) {
		s.Service = authService
	}
}
