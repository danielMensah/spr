package sql

import (
	"database/sql"
	"fmt"
)

type SQL struct {
	sql *sql.DB
}

func New(dbType string, connStr string) (SQL, error) {
	db, err := sql.Open(dbType, connStr)

	if err != nil {
		return SQL{}, fmt.Errorf("opening connection: %w", err)
	}

	s := SQL{
		sql: db,
	}

	return s, nil
}

func (s *SQL) Exec(script string) error {
	_, err := s.sql.Exec(script)

	if err != nil {
		return fmt.Errorf("executing script: %w", err)
	}

	return nil
}

func (s *SQL) Close() error {
	return s.sql.Close()
}
