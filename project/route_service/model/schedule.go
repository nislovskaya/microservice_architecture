package model

import "time"

type Schedule struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	RouteID   uint      `json:"routeId" gorm:"not null"`
	StartTime time.Time `json:"startTime" gorm:"not null"`
	EndTime   time.Time `json:"endTime" gorm:"not null"`
	Available int       `json:"available" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
