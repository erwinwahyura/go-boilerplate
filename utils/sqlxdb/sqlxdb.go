package sqlxdb

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	zlog "github.com/rs/zerolog/log"
)

type (
	// SqlxConfig mongo config
	SqlxConfig struct {
		Driver, Host, Port, Username, Password, DatabaseName string
	}
)

// NewSqlxDB sql db
func NewSqlxDsn(driver, dsn string) *sqlx.DB {

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		zlog.Fatal().Msgf("error when sqlx.Connect, error: %v", err.Error())
	}

	// Setup Connection
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Printf("ping %s", driver)
	if err := db.Ping(); err != nil {
		zlog.Fatal().Msgf("error when ping %s, error: %v", driver, err.Error())
	}

	return db
}
