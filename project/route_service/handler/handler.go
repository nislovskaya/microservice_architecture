package handler

import (
	"github.com/gorilla/mux"
	"github.com/nislovskaya/microservice_architecture/project/route_service/service"
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

	api := router.PathPrefix("/route").Subrouter()

	api.HandleFunc("/health", h.CheckHealth).Methods("GET")

	api.HandleFunc("/create", h.CreateRoute).Methods("POST")
	api.HandleFunc("/search", h.SearchRoutes).Methods("GET")
	api.HandleFunc("/{routeId}", h.GetRoute).Methods("GET")
	api.HandleFunc("/{routeId}", h.UpdateRoute).Methods("PUT")
	api.HandleFunc("/{routeId}", h.DeleteRoute).Methods("DELETE")

	return router
}
