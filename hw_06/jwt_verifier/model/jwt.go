package model

import "github.com/golang-jwt/jwt/v5"

type Jwt struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
