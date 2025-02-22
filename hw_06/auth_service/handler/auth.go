package handler

import (
	"errors"
	"fmt"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/httperrors"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/model"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/tools/response"
	"net/http"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	email, password, err := getCredentials(r)
	if err != nil {
		resp.BadRequest(err.Error())
		return
	}

	if err = h.Service.Register(email, password); err != nil {
		h.Logger.Errorf("Failed to register user, error: %v", err)
		var badReqErr *httperrors.BadRequestError
		switch {
		case errors.As(err, &badReqErr):
			resp.BadRequest(badReqErr.Error())
			return
		default:
			resp.InternalServerError(err.Error())
			return
		}
	}

	message := fmt.Sprintf("User with email '%s' has successfully registered", email)

	h.Logger.Infof(message)

	resp.Created(&model.Message{Message: message})
}

func getCredentials(r *http.Request) (string, string, error) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse multipart form: %v", err)
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	return email, password, nil
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	email, password, err := getCredentials(r)
	if err != nil {
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

	resp.Ok(&model.Token{Token: accessToken})
}
