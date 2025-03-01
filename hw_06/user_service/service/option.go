package service

import "github.com/nislovskaya/microservice_architecture/hw_06/user_service/service/user"

type Option func(fs *Service)

func WithUserService(userService user.Service) Option {
	return func(s *Service) {
		s.Service = userService
	}
}
