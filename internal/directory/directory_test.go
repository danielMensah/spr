package directory

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		expectedErr string
	}{
		{
			name:        "ok",
			path:        "./test",
			expectedErr: "",
		},
		{
			name:        "invalid path",
			path:        "",
			expectedErr: "creating folders",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateFolder(tt.path)

			if tt.expectedErr == "" {
				_, err := os.Stat(tt.path)

				assert.NoError(t, err)
				assert.False(t, os.IsNotExist(err))

				err = os.Remove(tt.path)
				if err != nil {
					log.Fatal("could not remove csv test file: %w", err)
				}
			} else {
				assert.Contains(t, err.Error(), tt.expectedErr)
			}
		})
	}
}

func TestCopyToDestination(t *testing.T) {
	tests := []struct {
		name        string
		source      string
		destination string
		expectedErr string
	}{
		{
			name:        "can copy to destination",
			source:      "./schema",
			destination: "./current/schema",
			expectedErr: "",
		},
		{
			name:        "invalid source",
			source:      "",
			destination: "./current/schema",
			expectedErr: "get source file content",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateFolder(tt.source)
			assert.NoError(t, err)

			err = CreateFolder(tt.destination)
			assert.NoError(t, err)

			filePath := filepath.Join(tt.source, "test.sql")
			file, err := os.Create(filePath)
			if err != nil {
				assert.NoError(t, err)
			}
			defer file.Close()

			err = CopyToDestination(filePath, filepath.Join(tt.destination, "test.sql"))
			if tt.expectedErr == "" {
				assert.NoError(t, err)

				_, err := os.Stat(filepath.Join(tt.destination, "test.sql"))

				assert.False(t, os.IsNotExist(err))
				assert.NoError(t, err)
			} else {
				assert.Contains(t, err.Error(), tt.expectedErr)
			}

			cleanup()
		})
	}
}

func cleanup() {
	_ = os.RemoveAll("./schema")
	_ = os.RemoveAll("./current")
}
