package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/mattismoel/konnekt/internal/server"
	"github.com/mattismoel/konnekt/internal/service"
	"github.com/mattismoel/konnekt/internal/storage/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

const (
	MAX_STARTUP_DURATION = 10 * time.Second
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), MAX_STARTUP_DURATION)
	defer cancel()

	flag.Parse()

	dbConnStr := flag.String("dbConnStr", "file:./local.db", "The database connection string")
	host := flag.String("host", "localhost", "The host of the web server")
	port := flag.Int("port", 3000, "The port of the web server")

	db, err := sql.Open("sqlite3", *dbConnStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	userRepo, err := sqlite.NewUserRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	authRepo, err := sqlite.NewAuthRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	eventRepo, err := sqlite.NewEventRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	artistRepo, err := sqlite.NewArtistRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	venueRepo, err := sqlite.NewVenueRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	authService, err := service.NewAuthService(userRepo, authRepo)
	if err != nil {
		log.Fatal(err)
	}

	eventService, err := service.NewEventService(eventRepo, artistRepo, venueRepo)

	srv, err := server.New(
		server.WithAddress(net.JoinHostPort(*host, strconv.Itoa(*port))),
		server.WithAuthService(authService),
		server.WithEventService(eventService),
	)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
