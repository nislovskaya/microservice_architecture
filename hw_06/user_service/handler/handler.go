package handler

import (
	"github.com/gorilla/mux"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Logger *logrus.Entry
	*service.Service
}

func New(opts ...Option) *Handler {
	handler := &Handler{}

	for _, option := range opts {
		option(handler)
	}

	return handler
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", h.CheckHealth)

	router.HandleFunc("/user/{userId}", h.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", h.UpdateUser).Methods("PUT")

	return router
}
