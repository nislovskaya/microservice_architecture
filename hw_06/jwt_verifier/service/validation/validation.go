package validation

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nislovskaya/microservice_architecture/hw_06/jwt_verifier/model"
	"github.com/sirupsen/logrus"
)

type Service interface {
	ValidateToken(tokenString string) (*model.JwtClaims, error)
}
type validation struct {
	logger    *logrus.Entry
	secretKey string
}

func New(opts ...Option) Service {
	service := &validation{}

	for _, option := range opts {
		option(service)
	}

	return service
}

func (v *validation) ValidateToken(tokenString string) (*model.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(v.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(*model.JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
