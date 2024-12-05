package main

import (
	"encoding/json"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/health/", healthHandler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
