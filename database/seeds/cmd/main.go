package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/eschool/config/postgres"
	"github.com/umardev500/eschool/database/seeds"
)

func init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msgf("Error loading .env file: %s", err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	ctx := context.Background()
	db := postgres.NewPostgres(ctx)
	trx := postgres.NewTransaction(db)

	seeder := seeds.NewSeeder(trx)
	err := trx.WithTransaction(ctx, func(ctx context.Context) error {
		return seeder.Creds()
	})

	if err != nil {
		log.Fatal().Msgf("Error seeding database: %s", err)
	}

	log.Info().Msgf("Done seeding database")
}
