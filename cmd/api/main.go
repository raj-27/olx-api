package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/raj-27/olx-api/internal/config"
	"github.com/raj-27/olx-api/internal/handlers"
)

func main() {

	cfg := config.MustLoad()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlers.Health())

	server := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 10,
	}

	slog.Info("server started", slog.String("addr", fmt.Sprintf("localhost:%v", cfg.Port)))

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
