package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewPostgres(ctx context.Context) *sqlx.DB {
	start := time.Now()
	log.Info().Msg("Connecting to Postgres 🛠️...")
	db, _ := sqlx.ConnectContext(ctx, "postgres", os.Getenv("DATABASE_URL"))
	err := db.Ping()
	if err != nil {
		log.Fatal().Msgf("Error connecting to database: %s", err)
	}

	duration := time.Since(start)
	msg := fmt.Sprintf("Connected to Postgres \033[32m🎉 (\U000023F3 %s)\033[0m", duration)
	log.Info().Msg(msg)

	return db
}
