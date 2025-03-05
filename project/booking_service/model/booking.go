package model

import "time"

type Booking struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"userId"`
	RouteID   uint      `json:"routeId"`
	Seats     int       `json:"seats"`
	CreatedAt time.Time `json:"createdAt"`
}
