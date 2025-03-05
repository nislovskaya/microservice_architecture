package model

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
