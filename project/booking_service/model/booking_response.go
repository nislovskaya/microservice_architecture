package model

type BookingResponse struct {
	Booking Booking `json:"booking"`
	Route   Route   `json:"route"`
}
