package model

import "github.com/golang-jwt/jwt/v5"

type Jwt struct {
	UserID uint `json:"userID"`
	jwt.RegisteredClaims
}
