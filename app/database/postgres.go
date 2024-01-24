package database

import (
	"github.com/erwinwahyura/go-boilerplate/app/model"
	"github.com/erwinwahyura/go-boilerplate/utils/sqlxdb"
	"github.com/jmoiron/sqlx"
)

var (
	DRIVER_POSTGRES = "postgres"
)

type PostgresCollection struct {
	Master *sqlx.DB
	Slave  *sqlx.DB
}

func NewPostgresCollection(config model.Config) PostgresCollection {

	// DB
	master := sqlxdb.NewSqlxDsn(DRIVER_POSTGRES, config.Database.DB.Master)
	slave := sqlxdb.NewSqlxDsn(DRIVER_POSTGRES, config.Database.DB.Slave)

	return PostgresCollection{
		Master: master,
		Slave:  slave,
	}
}
