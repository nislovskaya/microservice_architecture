package model

import "time"

type Booking struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"userId" gorm:"not null"`
	RouteID   uint      `json:"routeId" gorm:"not null"`
	Seats     int       `json:"seats" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null;default:'pending'"` // pending, confirmed, cancelled
	BookedAt  time.Time `json:"bookedAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
