package main

import (
	"net/http"
)

// addRoutes()
func addRoutes(mux *http.ServeMux, config *Config) {
	_ = config
	mux.Handle("GET /v1/readiness", handleReadiness())
	mux.Handle("GET /v1/err", handleError())
}

func handleReadiness() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type StatusMsg struct {
			Status string `json:"status"`
		}
		respondJSON(w, http.StatusOK, StatusMsg{
			Status: http.StatusText(http.StatusOK),
		})
	})
}

func handleError() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	})
}