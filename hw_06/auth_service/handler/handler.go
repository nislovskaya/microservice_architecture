package handler

import (
	"github.com/gorilla/mux"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/cmd/middleware"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/service"
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

	api := router.PathPrefix("/auth").Subrouter()

	api.HandleFunc("/health", h.CheckHealth).Methods("GET")
	api.HandleFunc("/register", h.Register).Methods("POST")
	api.HandleFunc("/login", h.Login).Methods("POST")
	api.HandleFunc("/validate", h.ValidateToken).Methods("POST")

	protected := api.PathPrefix("").Subrouter()
	mdw := middleware.New(h.Logger, h.Service)
	protected.Use(mdw.AuthMiddleware())

	protected.HandleFunc("/logout", h.Logout).Methods("POST")

	return router
}
