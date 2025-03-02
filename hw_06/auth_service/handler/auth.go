package handler

import (
	"errors"
	"fmt"
	"github.com/nislovskaya/golang_tools/response"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/httperrors"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/model"
	"net/http"
	"strconv"
	"strings"
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

func (h *Handler) ValidateToken(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	token := r.Header.Get("Authorization")
	if token == "" {
		h.Logger.Error("Missing Authorization header")
		resp.Unauthorized("Missing Authorization header")
		return
	}

	token = strings.Replace(token, "Bearer ", "", 1)
	if err := h.Service.Logout(token); err != nil {
		h.Logger.Errorf("Failed to logout: %v", err)
		resp.InternalServerError(err.Error())
		return
	}

	h.Logger.Info("User logged out successfully")
	resp.Ok(&model.Message{
		Message: "Logged out successfully",
	})
}

func getCredentials(r *http.Request) (string, string, error) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse multipart form: %v", err)
	}

	return r.FormValue("email"), r.FormValue("password"), nil
}
