package model

import "time"

type Route struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FromCity  string    `json:"fromCity" gorm:"column:from_city;not null"`
	ToCity    string    `json:"toCity" gorm:"column:to_city;not null"`
	Departure time.Time `json:"departure" gorm:"not null"`
	Arrival   time.Time `json:"arrival" gorm:"not null"`
	Capacity  int       `json:"capacity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}
