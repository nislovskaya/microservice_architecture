package model

import "time"

type Booking struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"userId" gorm:"not null"`
	RouteID   uint      `json:"routeId" gorm:"not null"`
	Seats     int       `json:"seats" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
}
