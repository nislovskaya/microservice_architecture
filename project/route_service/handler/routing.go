package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nislovskaya/golang_tools/response"
	"github.com/nislovskaya/microservice_architecture/project/route_service/model"
	"net/http"
	"strconv"
)

func (h *Handler) CreateRoute(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	var route model.Route
	if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
		h.Logger.Errorf("Error decoding request body: %v", err)
		resp.BadRequest(err.Error())
		return
	}

	if err := h.Service.CreateRoute(&route); err != nil {
		h.Logger.Errorf("Error creating route: %v", err)
		resp.InternalServerError(err.Error())
		return
	}

	resp.Created(&route)
}

func (h *Handler) GetRoute(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["routeId"], 10, 32)
	if err != nil {
		h.Logger.Errorf("Error parsing route ID: %v", err)
		resp.BadRequest("invalid route ID")
		return
	}

	route, err := h.Service.GetRoute(uint(id))
	if err != nil {
		h.Logger.Errorf("Error getting route: %v", err)
		resp.InternalServerError(err.Error())
		return
	}

	resp.Ok(route)
}

func (h *Handler) UpdateRoute(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["routeId"], 10, 32)
	if err != nil {
		h.Logger.Errorf("Error parsing route ID: %v", err)
		resp.BadRequest("invalid route ID")
		return
	}

	var route model.Route
	if err = json.NewDecoder(r.Body).Decode(&route); err != nil {
		h.Logger.Errorf("Error decoding request body: %v", err)
		resp.BadRequest(err.Error())
		return
	}

	route.ID = uint(id)
	if err = h.Service.UpdateRoute(&route); err != nil {
		h.Logger.Errorf("Error updating route: %v", err)
		resp.InternalServerError(err.Error())
		return
	}

	resp.Ok(&route)
}

func (h *Handler) DeleteRoute(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["routeId"], 10, 32)
	if err != nil {
		h.Logger.Errorf("Error parsing route ID: %v", err)
		resp.BadRequest("invalid route ID")
		return
	}

	if err = h.Service.DeleteRoute(uint(id)); err != nil {
		h.Logger.Errorf("Error deleting route: %v", err)
		resp.InternalServerError(err.Error())
		return
	}

	resp.NoContent()
}

func (h *Handler) SearchRoutes(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	date := r.URL.Query().Get("departure")

	if from == "" || to == "" {
		resp.BadRequest("from and to parameters are required")
		return
	}

	routes, err := h.Service.SearchRoutes(from, to, date)
	if err != nil {
		h.Logger.Errorf("Error searching routes: %v", err)
		resp.InternalServerError(err.Error())
		return
	}

	resp.Ok(routes)
}
