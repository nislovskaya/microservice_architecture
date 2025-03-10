package handler

import (
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/service"
	"github.com/sirupsen/logrus"
)

type Option func(fs *Handler)

func WithLogger(logger *logrus.Entry) Option {
	return func(h *Handler) {
		h.Logger = logger
	}
}

func WithService(service *service.Service) Option {
	return func(h *Handler) {
		h.Service = service
	}
}
