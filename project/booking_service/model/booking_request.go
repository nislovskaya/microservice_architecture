package model

type BookingRequest struct {
	RouteID uint `json:"routeId"`
	Seats   int  `json:"seats"`
}
