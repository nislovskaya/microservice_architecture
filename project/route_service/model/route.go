package model

import "time"

type Route struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FromCity  string    `json:"fromCity" gorm:"not null"`
	ToCity    string    `json:"toCity" gorm:"not null"`
	Departure int       `json:"departure" gorm:"not null"`
	Arrival   int       `json:"arrival" gorm:"not null"`
	Capacity  int       `json:"capacity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
}
