package main

import (
	"flag"
	"log"

	"github.com/mattismoel/konnekt/rest"
	"github.com/mattismoel/konnekt/sqlite"
)

func main() {
	dsn := flag.String("dsn", "file:local.db", "The DSN for the database")
	host := flag.String("host", "localhost", "Host of the server")
	port := flag.Int("port", 3000, "The port of the server")

	repo := sqlite.New(*dsn)

	err := repo.Open()
	if err != nil {
		log.Fatal(err)
	}

	eventService := sqlite.NewEventService(repo)
	userService := sqlite.NewUserService(repo)
	genreService := sqlite.NewGenreService(repo)

	srv, err := rest.NewServer(rest.Cfg{
		EventService: eventService,
		UserService:  userService,
		GenreService: genreService,
		Host:         *host,
		Port:         *port,
	})

	if err != nil {
		log.Fatal(err)
	}

	if err = srv.Start(); err != nil {
		log.Fatal(err)
	}
}
