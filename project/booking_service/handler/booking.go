package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/nislovskaya/golang_tools/response"
	"github.com/nislovskaya/microservice_architecture/project/booking_service/model"
)

func (h *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	var bookingReq model.BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&bookingReq); err != nil {
		h.Logger.Errorf("Error decoding request body: %v", err)
		resp.BadRequest(err.Error())
		return
	}

	userID := r.Context().Value("userID").(uint)

	booking := &model.Booking{
		UserID:    userID,
		RouteID:   bookingReq.RouteID,
		Seats:     bookingReq.Seats,
		CreatedAt: time.Now(),
	}

	if err := h.Service.CreateBooking(booking); err != nil {
		h.Logger.Errorf("Error creating booking: %v", err)
		resp.InternalServerError(err.Error())
		return
	}

	resp.Created(booking)
}

func (h *Handler) GetBooking(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["bookingId"], 10, 32)
	if err != nil {
		h.Logger.Errorf("Error parsing booking ID: %v", err)
		resp.BadRequest("invalid booking ID")
		return
	}

	booking, err := h.Service.GetBooking(uint(id))
	if err != nil {
		h.Logger.Errorf("Error getting booking: %v", err)
		resp.InternalServerError(err.Error())
		return
	}

	resp.Ok(booking)
}

func (h *Handler) GetUserBookings(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	userID := r.Context().Value("userID").(uint)

	bookings, err := h.Service.GetUserBookings(userID)
	if err != nil {
		h.Logger.Errorf("Error getting user bookings: %v", err)
		resp.InternalServerError(err.Error())
		return
	}

	resp.Ok(bookings)
}

func (h *Handler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["bookingId"], 10, 32)
	if err != nil {
		h.Logger.Errorf("Error parsing booking ID: %v", err)
		resp.BadRequest("invalid booking ID")
		return
	}

	if err = h.Service.CancelBooking(uint(id)); err != nil {
		h.Logger.Errorf("Error cancelling booking: %v", err)
		resp.InternalServerError(err.Error())
		return
	}

	resp.NoContent()
}
