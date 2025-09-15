package main

import (
	"context"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mattismoel/konnekt/backend/internal/psql"
	"github.com/mattismoel/konnekt/backend/internal/server"
)

const (
	MAX_STARTUP_DURATION = 10 * time.Second
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), MAX_STARTUP_DURATION)
	defer cancel()

	connStr := os.Getenv("DATABASE_URL")
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		slog.Error("Could not parse server port", "error", err.Error())
		return
	}

	dbPool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		slog.Error("Could not connect to database", "error", err.Error())
		return
	}

	defer dbPool.Close()

	artistRepo := psql.ArtistRepo{Pool: dbPool}

	server := server.Server{
		ArtistRepo:  artistRepo,
	}

	if err := server.Start(host, port); err != nil {
		slog.Error("Could not start server", "error", err.Error())
		return
	}
}
