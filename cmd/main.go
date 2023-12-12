package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/eschool/application"
	"github.com/umardev500/eschool/config/database"
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
	database.NewPostgres(context.Background())
	app := application.NewApp()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	// Start the application
	err := app.Start(ctx)
	if err != nil {
		log.Fatal().Msgf("Error starting application: %s", err)
	}
}
