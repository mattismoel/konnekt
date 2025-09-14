package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/urfave/cli/v3"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "db",
				Usage: "database commands",
				Commands: []*cli.Command{
					{
						Name:    "initialise",
						Usage:   "initialises the database",
						Aliases: []string{"i"},
						Action: func(ctx context.Context, c *cli.Command) error {
							connStr := os.Getenv("DATABASE_URL")

							conn, err := pgx.Connect(ctx, connStr)
							if err != nil {
								return fmt.Errorf("Could not connect to database: %v", err)
							}

							if err := conn.Ping(ctx); err != nil {
								return fmt.Errorf("Could not ping database: %v", err)
							}

							if err := clearDatabase(ctx, conn); err != nil {
								return fmt.Errorf("Could not reset database: %v", err)
							}

							if err := initialiseDatabase(ctx, conn); err != nil {
								return fmt.Errorf("Could not initialise database: %v", err)
							}

							return nil
						},
					},
					{
						Name:    "seed",
						Aliases: []string{"s"},
						Usage:   "seeds the database",
						Action: func(ctx context.Context, c *cli.Command) error {
							dbConnStr := os.Getenv("DATABASE_URL")
							conn, err := pgx.Connect(ctx, dbConnStr)
							if err != nil {
								return fmt.Errorf("Could not connect to database: %v", err)
							}

							if err := conn.Ping(ctx); err != nil {
								return fmt.Errorf("Could not ping database: %v", err)
							}

							if err := seedDb(ctx, conn); err != nil {
								return fmt.Errorf("Could not seed database: %v", err)
							}

							return nil
						},
					},
				},
			},
		},
	}

	if err := cmd.Run(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}
