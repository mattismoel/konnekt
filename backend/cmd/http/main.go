package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
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
	MAX_STARTUP_DURATION = 10 * time.Second
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
		log.Fatal(err)
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	memberRepo, err := sqlite.NewMemberRepository(db)
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

	teamRepo, err := sqlite.NewTeamRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	authService, err := service.NewAuthService(memberRepo, authRepo, teamRepo)
	if err != nil {
		log.Fatal(err)
	}

	s3Store, err := s3.NewS3ObjectStore(*s3Region, *s3Bucket)
	if err != nil {
		log.Fatal(err)
	}

	memberService, err := service.NewMemberService(memberRepo, teamRepo, s3Store)
	if err != nil {
		log.Fatal(err)
	}

	eventService, err := service.NewEventService(eventRepo, artistRepo, venueRepo, s3Store)
	if err != nil {
		log.Fatal(err)
	}

	artistService, err := service.NewArtistService(artistRepo, eventRepo, s3Store)
	if err != nil {
		log.Fatal(err)
	}

	venueService := service.NewVenueService(venueRepo)

	teamService := service.NewTeamService(teamRepo, memberRepo, authRepo)

	srv, err := server.New(
		server.WithTeamService(teamService),
		server.WithAddress(net.JoinHostPort(*host, strconv.Itoa(*port))),
		server.WithCORSOrigins(*origin),
		server.WithAuthService(authService),
		server.WithMemberService(memberService),
		server.WithEventService(eventService),
		server.WithArtistService(artistService),
		server.WithVenueService(venueService),
	)

	slog.Info("Started server", "host", *host, "port", *port)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
