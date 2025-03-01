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

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		h.Logger.Errorf("Failed to decode body, error: %v", err)
		resp.BadRequest(err.Error())
		return
	}

	user, err := h.Service.GetUserByID(uint(id))
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

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		h.Logger.Errorf("Failed to decode body, error: %v", err)
		resp.BadRequest(err.Error())
		return
	}

	var user model.User
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.Logger.Errorf("Failed to decode body, error: %v", err)
		resp.BadRequest("Invalid body")
		return
	}

	user.ID = uint(id)

	if err = h.Service.UpdateUser(&user); err != nil {
		h.Logger.Errorf("Failed to update user, error: %v", err)
		resp.InternalServerError("Failed to update user")
		return
	}

	h.Logger.Infof("Updated user with ID: %d", user.ID)
	resp.Ok(user)
}
