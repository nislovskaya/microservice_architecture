package handler

import (
	"errors"
	"fmt"
	"github.com/nislovskaya/golang_tools/response"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/httperrors"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/model"
	"net/http"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	email, password, err := getCredentials(r)
	if err != nil {
		h.Logger.Errorf("Error getting credentials: %v", err)
		resp.BadRequest(err.Error())
		return
	}

	userID, err := h.Service.Register(email, password)
	if err != nil {
		h.Logger.Errorf("Error registering user: %v", err)
		var conflictError *httperrors.ConflictError
		switch {
		case errors.As(err, &conflictError):
			resp.Conflict(conflictError.Error())
			return
		default:
			resp.InternalServerError(err.Error())
			return
		}
	}

	h.Logger.Infof(fmt.Sprintf("User '%s' registered, ID: %d", email, userID))
	resp.Created(&model.User{
		ID:    userID,
		Email: email,
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	email, password, err := getCredentials(r)
	if err != nil {
		h.Logger.Errorf("Error getting credentials: %v", err)
		resp.BadRequest(err.Error())
		return
	}

	accessToken, err := h.Service.Login(email, password)
	if err != nil {
		h.Logger.Errorf("Failed to login user, error: %v", err)
		var unauthorizedError *httperrors.UnauthorizedError
		switch {
		case errors.As(err, &unauthorizedError):
			resp.Unauthorized(err.Error())
			return
		default:
			resp.InternalServerError(err.Error())
			return
		}
	}

	h.Logger.Infof("User with email '%s' logged in", email)
	resp.Ok(&model.Token{
		Token: accessToken,
	})
}

func getCredentials(r *http.Request) (string, string, error) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse multipart form: %v", err)
	}

	return r.FormValue("email"), r.FormValue("password"), nil
}
