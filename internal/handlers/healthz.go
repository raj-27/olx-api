package handlers

import (
	"log/slog"
	"net/http"
)

// Health returns the current API health status.
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write([]byte(`{"status":"ok"}`)); err != nil {
			slog.Error("health response failed", slog.Any("error", err))
		}
	}
}
