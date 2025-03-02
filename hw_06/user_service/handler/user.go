package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nislovskaya/golang_tools/response"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/model"
	"net/http"
	"strconv"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	currentUserID, err := strconv.ParseUint(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		h.Logger.Errorf("Failed to parse x-user-id header, error: %v", err)
		resp.Unauthorized("Invalid user ID")
		return
	}

	params := mux.Vars(r)
	requestedUserID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		h.Logger.Errorf("Failed to parse userId, error: %v", err)
		resp.BadRequest(err.Error())
		return
	}

	if requestedUserID != currentUserID {
		h.Logger.Errorf("User %d attempted to access profile of user %d", currentUserID, requestedUserID)
		resp.Forbidden("Access denied")
		return
	}

	user, err := h.Service.GetUserByID(uint(requestedUserID))
	if err != nil {
		h.Logger.Errorf("Failed to get user, error: %v", err)
		resp.NotFound("User not found")
		return
	}

	h.Logger.Infof("Getting user with ID: %d", user.ID)
	resp.Ok(user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	currentUserID, err := strconv.ParseUint(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		h.Logger.Errorf("Failed to parse x-user-id header, error: %v", err)
		resp.Unauthorized("Invalid user ID")
		return
	}

	params := mux.Vars(r)
	requestedUserID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		h.Logger.Errorf("Failed to parse userId, error: %v", err)
		resp.BadRequest(err.Error())
		return
	}

	if requestedUserID != currentUserID {
		h.Logger.Errorf("User %d attempted to update profile of user %d", currentUserID, requestedUserID)
		resp.Forbidden("Access denied")
		return
	}

	var user model.User
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.Logger.Errorf("Failed to decode body, error: %v", err)
		resp.BadRequest("Invalid body")
		return
	}

	user.ID = uint(requestedUserID)

	if err = h.Service.UpdateUser(&user); err != nil {
		h.Logger.Errorf("Failed to update user, error: %v", err)
		resp.InternalServerError("Failed to update user")
		return
	}

	h.Logger.Infof("Updated user with ID: %d", user.ID)
	resp.Ok(user)
}
