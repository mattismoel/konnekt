package main

import (
	"flag"
	"log"

	"github.com/mattismoel/konnekt/internal/rest"
	"github.com/mattismoel/konnekt/internal/service"
	"github.com/mattismoel/konnekt/internal/storage/sqlite"
)

func main() {
	dsn := flag.String("dsn", "file:local.db", "The DSN for the database")
	host := flag.String("host", "localhost", "Host of the server")
	port := flag.Int("port", 3000, "The port of the server")

	store := sqlite.NewStore(*dsn)

	err := store.Open()
	if err != nil {
		log.Fatal(err)
	}

	eventRepo := sqlite.NewEventRepository(store)
	userRepo := sqlite.NewUserRepository(store)
	genreRepo := sqlite.NewGenreRepository(store)

	userService := service.NewUserService(userRepo)
	eventService := service.NewEventService(eventRepo)
	genreService := service.NewGenreService(genreRepo)

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
