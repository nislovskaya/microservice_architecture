package service

import "github.com/nislovskaya/microservice_architecture/hw_04/crud_app/service/user"

type Option func(fs *Service)

func WithUserService(userService user.Service) Option {
	return func(s *Service) {
		s.Service = userService
	}
}
