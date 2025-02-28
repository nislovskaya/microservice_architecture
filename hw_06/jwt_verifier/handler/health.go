package handler

import (
	"github.com/nislovskaya/golang_tools/response"
	"net/http"
)

func (h *Handler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w, h.Logger)

	message := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	h.Logger.Infof("Health checked with status: %s", message.Status)
	resp.Ok(message)
}
