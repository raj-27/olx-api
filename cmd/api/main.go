package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/raj-27/olx-api/internal/config"
	"github.com/raj-27/olx-api/internal/db"
	"github.com/raj-27/olx-api/internal/handlers"
)

func main() {
	// Load application configuration.
	cfg := config.MustLoad()

	// Establish database connection.
	_, err := db.Connect(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	slog.Info("database connected")

	// Register HTTP routes.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health())

	// Configure HTTP server with timeouts.
	server := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	slog.Info("server started", slog.String("addr", fmt.Sprintf("localhost:%s", cfg.Port)))

	// Start the HTTP server.
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
