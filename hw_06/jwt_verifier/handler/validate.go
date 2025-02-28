package handler

import (
	"fmt"
	"github.com/nislovskaya/golang_tools/response"
	"github.com/nislovskaya/microservice_architecture/hw_06/jwt_verifier/model"
	"net/http"
	"strings"
)

func (h *Handler) Validate(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	token := r.Header.Get("Authorization")
	if token == "" {
		resp.Unauthorized("Missing Authorization header")
		return
	}

	token = strings.Replace(token, "Bearer ", "", 1)
	_, err := h.Service.ValidateToken(token)
	if err != nil {
		resp.Unauthorized(err.Error())
		return
	}

	message := fmt.Sprintf("Token is valid")

	h.Logger.Infof(message)

	resp.Ok(&model.Message{Message: message})
}
