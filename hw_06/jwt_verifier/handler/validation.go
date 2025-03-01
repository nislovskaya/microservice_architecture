package handler

import (
	"github.com/nislovskaya/golang_tools/response"
	"github.com/nislovskaya/microservice_architecture/hw_06/jwt_verifier/model"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) Validate(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	token := r.Header.Get("Authorization")
	if token == "" {
		h.Logger.Error("Missing Authorization header")
		resp.Unauthorized("Missing Authorization header")
		return
	}

	token = strings.Replace(token, "Bearer ", "", 1)
	claims, err := h.Service.ValidateToken(token)
	if err != nil {
		h.Logger.Errorf("Error validating token: %v", err)
		resp.Unauthorized(err.Error())
		return
	}

	w.Header().Set("x-user-id", strconv.Itoa(int(claims.UserID)))

	h.Logger.Infof("Token validated for user: %d", claims.UserID)
	resp.Ok(&model.Message{
		Message: "Token is valid",
	})
}
