package config

import (
	"fmt"
)

// Postgres holds the data struct in order to create a connection
type Postgres struct {
	Username       string `yaml:"username,omitempty"`
	Password       string `yaml:"password,omitempty"`
	Database       string `yaml:"database,omitempty"`
	Address        string `yaml:"address,omitempty"`
	SSLMode        string `yaml:"sslMode,omitempty"`
	PoolMaxConns   string `yaml:"poolMaxConns,omitempty"`
	ConnectTimeout string `yaml:"connectTimeout,omitempty"`
}

// Database is used to simply define the methods that can be used
type Database interface {
	Config() Postgres
}

// Config returns the entire Postgres struct in case the user want to do something specific with the data
func (c Config) Config() Postgres {
	return c.Database
}

// ConnectionString is a helper func that returns the connection string for the db
func (p Postgres) ConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s&connect_timeout=%s",
		p.Username, p.Password, p.Address, p.Database, p.SSLMode, p.ConnectTimeout,
	)
}

// PoolConnectionString is a helper func that returns the connection string for the db with additional pgxpool variables
func (p Postgres) PoolConnectionString() string {
	return fmt.Sprintf(
		"%s&pool_max_conns=%s",
		p.ConnectionString(), p.PoolMaxConns,
	)
}
