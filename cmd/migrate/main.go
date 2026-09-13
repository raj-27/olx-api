package main

import (
	"errors"
	"log"
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/raj-27/olx-api/internal/config"
)

func main() {
	// Make sure a migration command was provided.
	// Usage: make migrate up
	//        make migrate down
	if len(os.Args) < 2 {
		log.Fatal("usage: make migrate <up | down>")
	}

	// Load application configuration.
	cfg := config.MustLoad()

	// Create a migration instance using the local migrations directory
	// and the PostgreSQL database URL from configuration.
	m, err := migrate.New("file://migrations", cfg.DATABASE_URL)
	if err != nil {
		slog.Error("failed to initialize migration", slog.Any("error", err))
		return
	}

	// Always close the migration instance before exiting.
	defer func() {
		sourceErr, databaseErr := m.Close()

		if sourceErr != nil {
			slog.Error("failed to close migration source", slog.Any("error", sourceErr))
		}

		if databaseErr != nil {
			slog.Error("failed to close migration database", slog.Any("error", databaseErr))
		}
	}()

	// Execute the requested migration command.
	switch os.Args[1] {

	case "up":
		slog.Info("running database migrations", slog.String("direction", "up"))

		if err := m.Up(); err != nil {

			// ErrNoChange means all migrations are already applied.
			if errors.Is(err, migrate.ErrNoChange) {
				slog.Info("database is already up to date")
				return
			}

			slog.Error("failed to run migrations", slog.String("direction", "up"), slog.Any("error", err))
			return
		}
		slog.Info("database migrations completed successfully", slog.String("direction", "up"))

	case "down":
		slog.Info("rolling back database migrations", slog.String("direction", "down"))
		if err := m.Down(); err != nil {

			// ErrNoChange means there are no migrations to roll back.
			if errors.Is(err, migrate.ErrNoChange) {
				slog.Info("no migrations to roll back")
				return
			}
			slog.Error("failed to rollback migrations", slog.String("direction", "down"), slog.Any("error", err))
			return
		}
		slog.Info("database migrations rolled back successfully", slog.String("direction", "down"))
	default:
		slog.Error("unknown migration command", slog.String("command", os.Args[1]))
		os.Exit(1)
	}
}
