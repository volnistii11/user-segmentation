package database

import "github.com/jmoiron/sqlx"

type Storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}
