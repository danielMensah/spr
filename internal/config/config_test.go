package config

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		setPath     bool
		setName     bool
		config      *bytes.Buffer
		expected    Interface
		expectedErr string
	}{
		{
			name:        "read config err",
			path:        "",
			setPath:     true,
			setName:     true,
			config:      bytes.NewBufferString(``),
			expected:    nil,
			expectedErr: "reading config",
		},
		{
			name:        "parse error",
			path:        ".",
			setPath:     true,
			setName:     true,
			config:      bytes.NewBufferString(`{`),
			expected:    nil,
			expectedErr: "reading config",
		},
		{
			name:    "new config",
			path:    ".",
			setPath: true,
			setName: true,
			config: bytes.NewBufferString(`stores:
  store_1:
    name: "name"
    apiKey: "apiKey"
    sharedSecret: "sharedSecret"
    passwordToken: "passwordToken"
    shopifyURL: "shopifyURL"
    index: "index"
database:
  username: "username"
  password: "password"
  address: "address"
  database: "database"`),
			expected: Config{
				Database: Postgres{
					Username:       "username",
					Password:       "password",
					Database:       "database",
					Address:        "address",
					PoolMaxConns:   "10",
					ConnectTimeout: "10",
				},
			},
			expectedErr: "",
		},
		{
			name:        "missing path env",
			path:        "",
			setPath:     false,
			setName:     true,
			config:      bytes.NewBufferString(``),
			expected:    nil,
			expectedErr: "missing CONFIG_PATH",
		},
		{
			name:        "missing name env",
			path:        "",
			setPath:     true,
			setName:     false,
			config:      bytes.NewBufferString(``),
			expected:    nil,
			expectedErr: "missing CONFIG_NAME",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := os.WriteFile("store-sync-temp.yaml", tt.config.Bytes(), os.ModePerm); err != nil {
				assert.Nil(t, err)
			}

			if tt.setPath {
				_ = os.Setenv(envConfigPath, tt.path)
			}

			if tt.setName {
				_ = os.Setenv(envConfigName, "store-sync-temp")
			}

			got, err := New()

			t.Logf("%T", err)

			assert.Equal(t, tt.expected, got)

			if tt.expectedErr == "" {
				assert.NoError(t, err)
			} else {
				assert.Contains(t, err.Error(), tt.expectedErr)
			}

			if err := os.Remove("store-sync-temp.yaml"); err != nil {
				assert.Nil(t, err)
			}

			os.Clearenv()
		})
	}
}
