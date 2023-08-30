package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewConnection(driver string, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a db conn")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to ping the db")
	}

	return db, nil
}
