package middleware

import (
	"github.com/nislovskaya/golang_tools/response"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/service/auth"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

type Middleware struct {
	logger  *logrus.Entry
	service auth.Service
}

func New(logger *logrus.Entry, service auth.Service) *Middleware {
	return &Middleware{
		logger:  logger,
		service: service,
	}
}

func (m *Middleware) AuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resp := response.New(w, m.logger)

			token := r.Header.Get("Authorization")
			if token == "" {
				m.logger.Error("No token provided")
				resp.Unauthorized("Missing Authorization header")
				return
			}

			token = strings.Replace(token, "Bearer ", "", 1)
			claims, err := m.service.ValidateToken(token)
			if err != nil {
				m.logger.Errorf("Error validating token: %s", err.Error())
				resp.Unauthorized(err.Error())
				return
			}

			r.Header.Set("x-user-id", strconv.Itoa(int(claims.UserID)))
			next.ServeHTTP(w, r)
		})
	}
}
