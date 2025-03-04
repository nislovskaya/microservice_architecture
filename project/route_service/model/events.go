package model

const (
	RouteCreated     = "ROUTE_CREATED"
	RouteUpdated     = "ROUTE_UPDATED"
	RouteDeleted     = "ROUTE_DELETED"
	BookingCreated   = "BOOKING_CREATED"
	BookingCanceled  = "BOOKING_CANCELED"
	BookingConfirmed = "BOOKING_CONFIRMED"
)

type RouteEvent struct {
	Type      string `json:"type"`
	RouteID   uint   `json:"routeId"`
	Capacity  int    `json:"capacity"`
	Timestamp string `json:"timestamp"`
}

type BookingEvent struct {
	Type      string `json:"type"`
	BookingID uint   `json:"bookingId"`
	RouteID   uint   `json:"routeId"`
	Seats     int    `json:"seats"`
	Timestamp string `json:"timestamp"`
}
