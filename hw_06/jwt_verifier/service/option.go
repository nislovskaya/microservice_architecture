package service

import (
	"github.com/nislovskaya/microservice_architecture/hw_06/jwt_verifier/service/validation"
)

type Option func(fs *Service)

func WithAuthService(validationService validation.Service) Option {
	return func(s *Service) {
		s.Service = validationService
	}
}
