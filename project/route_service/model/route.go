package model

import "time"

type Route struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null"`
	StartPoint string    `json:"startPoint" gorm:"not null"`
	EndPoint   string    `json:"endPoint" gorm:"not null"`
	Price      float64   `json:"price" gorm:"not null"`
	Capacity   int       `json:"capacity" gorm:"not null"`
	IsActive   bool      `json:"isActive" gorm:"default:true"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
