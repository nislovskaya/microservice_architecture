package model

type Route struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	StartPoint string  `json:"startPoint"`
	EndPoint   string  `json:"endPoint"`
	Price      float64 `json:"price"`
}
