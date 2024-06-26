package main

import "net/http"

func handleReadiness() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type statusMsg struct {
			Status string `json:"status"`
		}
		respondJSON(w, http.StatusOK, statusMsg{
			Status: http.StatusText(http.StatusOK),
		})
	})
}

func handleError() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	})
}
