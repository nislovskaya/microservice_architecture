package model

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserID uint `json:"userID"`
	jwt.RegisteredClaims
}
