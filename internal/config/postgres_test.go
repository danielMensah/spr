package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Config(t *testing.T) {
	tests := []struct {
		name     string
		database Postgres
		expected Postgres
	}{
		{
			name: "returns correctly",
			database: Postgres{
				Username: "username",
				Password: "password",
				Database: "database",
				Address:  "address",
			},
			expected: Postgres{
				Username: "username",
				Password: "password",
				Database: "database",
				Address:  "address",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				Database: tt.database,
			}
			got := c.Config()
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestPostgres_ConnectionString(t *testing.T) {
	tests := []struct {
		name     string
		database Postgres
		expected string
	}{
		{
			name: "returns correctly",
			database: Postgres{
				Username:       "username",
				Password:       "password",
				Database:       "database",
				Address:        "address",
				SSLMode:        "required",
				ConnectTimeout: "5",
			},
			expected: "postgres://username:password@address/database?sslmode=required&connect_timeout=5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.database.ConnectionString()

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestPostgres_PoolConnectionString(t *testing.T) {
	tests := []struct {
		name     string
		database Postgres
		expected string
	}{
		{
			name: "returns correctly",
			database: Postgres{
				Username:       "username",
				Password:       "password",
				Database:       "database",
				Address:        "address",
				SSLMode:        "required",
				ConnectTimeout: "5",
				PoolMaxConns:   "5",
			},
			expected: "postgres://username:password@address/database?sslmode=required&connect_timeout=5&pool_max_conns=5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.database.PoolConnectionString()

			assert.Equal(t, tt.expected, got)
		})
	}
}
