package handler

import (
	"github.com/gorilla/mux"
	"github.com/nislovskaya/microservice_architecture/project/booking_service/service"
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

	api := router.PathPrefix("/booking").Subrouter()

	api.HandleFunc("/health", h.CheckHealth).Methods("GET")
	api.HandleFunc("/create", h.CreateBooking).Methods("POST")
	api.HandleFunc("/{bookingId}", h.GetBooking).Methods("GET")
	api.HandleFunc("/user", h.GetUserBookings).Methods("GET")
	api.HandleFunc("/{bookingId}/cancel", h.CancelBooking).Methods("POST")

	return router
}
