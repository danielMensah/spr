package sql

import (
	"database/sql"
)

const NewPostgresError = "postgres initialization error"

func New(dbType string, connStr string) (*sql.DB, error) {
	db, err := sql.Open(dbType, connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}