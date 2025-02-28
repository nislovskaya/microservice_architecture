package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nislovskaya/golang_tools/response"
	"github.com/nislovskaya/microservice_architecture/hw_06/jwt_verifier/model"
)

func (h *Handler) Validate(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	token := r.Header.Get("Authorization")
	if token == "" {
		resp.Unauthorized("Missing Authorization header")
		return
	}

	token = strings.Replace(token, "Bearer ", "", 1)
	claims, err := h.Service.ValidateToken(token)
	if err != nil {
		resp.Unauthorized(err.Error())
		return
	}

	userID := fmt.Sprintf("%d", claims.UserID)
	w.Header().Set("x-user-id", userID)

	h.Logger.Infof("Token validated for user: %s", userID)
	resp.Ok(&model.Message{Message: "Token is valid"})
}
