package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/eschool/constant"
)

type PostgresQueries interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type PostgresDB interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

type Trx struct {
	db PostgresDB
}

type TxFn func(context.Context) error

// NewTransaction creates a new transaction.
//
// It takes a PostgresDB as a parameter and returns a pointer to a Trx.
func NewTransaction(db PostgresDB) *Trx {
	return &Trx{
		db: db,
	}
}

// WithTransaction executes a function inside a transaction.
//
// ctx: The context in which the transaction should be executed.
// fn: The function to be executed inside the transaction.
// err: The error returned by the function, if any.
func (t *Trx) WithTransaction(ctx context.Context, fn TxFn) (err error) {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			log.Error().Msgf("Rollback transaction: %s", err)
			tx.Rollback()
		} else {
			log.Info().Msgf("Commit transaction")
			err = tx.Commit()
		}
	}()

	ctx = context.WithValue(ctx, constant.CtxKeyTx, tx)
	err = fn(ctx)

	return
}

// NewPostgres creates a new connection to a Postgres database.
//
// ctx: The context.Context object used for cancellation and timeouts.
// Returns: A pointer to a sqlx.DB object representing the database connection.
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
