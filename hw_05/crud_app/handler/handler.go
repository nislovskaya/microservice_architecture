package handler

import (
	"github.com/gorilla/mux"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	router.Handle("/metrics", promhttp.Handler())

	router.HandleFunc("/health", h.CheckHealth)

	router.HandleFunc("/user", h.CreateUser).Methods("POST")
	router.HandleFunc("/user/{userId}", h.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", h.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", h.DeleteUser).Methods("DELETE")

	return router
}
