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

	// load config
	cfg := config.MustLoad()

	_, err := db.Connect(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("main.db.connect:%v", err)
	}

	slog.Info("db connected")

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
