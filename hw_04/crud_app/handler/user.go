package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/model"
	"net/http"
	"strconv"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Error{Code: 400, Message: "Invalid input"})
		h.Logger.Errorf("Failed to decode body, error: %v", err)
		return
	}

	if err := h.Service.CreateUser(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Error{Code: 500, Message: "Could not create user"})
		h.Logger.Errorf("Failed to create user, error: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	h.Logger.Infof("Created user with id = %d", user.ID)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["userId"], 10, 64)
	h.Logger.Debugf("user id: %v", id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Error{Code: 400, Message: "Invalid ID"})
		h.Logger.Errorf("Failed to decode body, error: %v", err)
		return
	}

	user, err := h.Service.GetUserByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Error{Code: 404, Message: "User not found"})
		h.Logger.Errorf("Failed to get user, error: %v", err)
		return
	}

	json.NewEncoder(w).Encode(user)
	h.Logger.Infof("Getting user with id = %d", user.ID)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["userId"], 10, 64)
	h.Logger.Debugf("user id: %v", id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Error{Code: 400, Message: "Invalid ID"})
		h.Logger.Errorf("Failed to decode body, error: %v", err)
		return
	}

	var user model.User
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Error{Code: 400, Message: "Invalid input"})
		h.Logger.Errorf("Failed to decode body, error: %v", err)
		return
	}

	user.ID = uint(id)

	if err = h.Service.UpdateUser(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Error{Code: 500, Message: "Could not update user"})
		h.Logger.Errorf("Failed to update user, error: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	h.Logger.Infof("Updated user with id = %d", user.ID)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["userId"], 10, 64)
	h.Logger.Debugf("user id: %v", id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Error{Code: 400, Message: "Invalid ID"})
		h.Logger.Errorf("Failed to decode body, error: %v", err)
		return
	}

	if err = h.Service.DeleteUser(uint(id)); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Error{Code: 404, Message: "User not found"})
		h.Logger.Errorf("Failed to delete user, error: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	h.Logger.Infof("Deleted user with id = %d", id)
}
