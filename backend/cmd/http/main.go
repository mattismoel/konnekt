package main

import (
	"context"
	"database/sql"
	"flag"
	"log/slog"
	"net"
	"strconv"
	"time"

	"github.com/mattismoel/konnekt/internal/object/s3"
	"github.com/mattismoel/konnekt/internal/server"
	"github.com/mattismoel/konnekt/internal/service"
	"github.com/mattismoel/konnekt/internal/storage/sqlite"
	_ "modernc.org/sqlite"
)

const (
	MAX_STARTUP_DURATION = 60 * time.Second
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), MAX_STARTUP_DURATION)
	defer cancel()

	dbConnStr := flag.String("dbConnStr", "file:./local.db", "The database connection string")
	origin := flag.String("origin", "http://localhost:4000", "The origin of the proxy web server")
	host := flag.String("host", "127.0.0.1", "The host of the web server")
	port := flag.Int("port", 8080, "The port of the web server")
	s3Region := flag.String("s3Region", "eu-north-1", "The region of the S3 bucket")
	s3Bucket := flag.String("s3Bucket", "konnekt-bucket", "The bucket name of the S3 bucket")

	flag.Parse()

	db, err := sql.Open("sqlite", *dbConnStr)
	if err != nil {
		slog.Error("Could not create database connection", "error", err)
		return
	}

	if err := db.PingContext(ctx); err != nil {
		slog.Error("Could not ping database", "error", err)
		return
	}

	contentRepo, err := sqlite.NewContentRepository(db)
	if err != nil {
		slog.Error("Could not create content repository", "error", err)
		return
	}

	memberRepo, err := sqlite.NewMemberRepository(db)
	if err != nil {
		slog.Error("Could not create member repository", "error", err)
		return
	}

	authRepo, err := sqlite.NewAuthRepository(db)
	if err != nil {
		slog.Error("Could not create auth repository", "error", err)
		return
	}

	eventRepo, err := sqlite.NewEventRepository(db)
	if err != nil {
		slog.Error("Could not create event repository", "error", err)
		return
	}

	artistRepo, err := sqlite.NewArtistRepository(db)
	if err != nil {
		slog.Error("Could not create artist repository", "error", err)
		return
	}

	venueRepo, err := sqlite.NewVenueRepository(db)
	if err != nil {
		slog.Error("Could not create venue repository", "error", err)
		return
	}

	teamRepo, err := sqlite.NewTeamRepository(db)
	if err != nil {
		slog.Error("Could not create team repository", "error", err)
		return
	}

	authService, err := service.NewAuthService(memberRepo, authRepo, teamRepo)
	if err != nil {
		slog.Error("Could not create auth service", "error", err)
		return
	}

	s3Store, err := s3.NewS3ObjectStore(ctx, *s3Region, *s3Bucket)
	if err != nil {
		slog.Error("Could not create S3 store", "error", err)
		return
	}

	memberService, err := service.NewMemberService(memberRepo, teamRepo, s3Store)
	if err != nil {
		slog.Error("Could not create member service", "error", err)
		return
	}

	eventService, err := service.NewEventService(eventRepo, artistRepo, venueRepo, s3Store)
	if err != nil {
		slog.Error("Could not create event service", "error", err)
		return
	}

	artistService, err := service.NewArtistService(artistRepo, eventRepo, s3Store)
	if err != nil {
		slog.Error("Could not create artist service", "error", err)
		return
	}

	venueService := service.NewVenueService(venueRepo)

	teamService := service.NewTeamService(teamRepo, memberRepo, authRepo)
	contentService := service.NewContentService(s3Store, contentRepo)

	srv, err := server.New(
		server.WithContentService(contentService),
		server.WithTeamService(teamService),
		server.WithAddress(net.JoinHostPort(*host, strconv.Itoa(*port))),
		server.WithCORSOrigins(*origin),
		server.WithAuthService(authService),
		server.WithMemberService(memberService),
		server.WithEventService(eventService),
		server.WithArtistService(artistService),
		server.WithVenueService(venueService),
	)

	if err != nil {
		slog.Error("Could not create server", "error", err)
		return
	}

	slog.Info("Started server", "host", *host, "port", *port, "origin", *origin)
	if err := srv.Start(); err != nil {
		slog.Error("Could not start server", "error", err)
		return
	}
}
