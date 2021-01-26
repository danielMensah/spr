package directory

import (
	"fmt"
	"github.com/danielMensah/sqlr/internal/params"
	"github.com/danielMensah/sqlr/pkg/errors"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

var absolutePath string

func init() {
	path, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	absolutePath = filepath.Join(path, "temp")
}

func TestBaseline(t *testing.T) {
	type args struct {
		path     string
		fileType string
		fileName string
	}
	tests := []struct {
		name            string
		args            args
		expectedMessage string
		expectedErr     error
	}{
		{
			name: "can baseline an sql file",
			args: args{
				path:     absolutePath,
				fileName: "example1.sql",
				fileType: params.FileTypes["-f"],
			},
			expectedMessage: BaselineResponse,
			expectedErr:     nil,
		},
		{
			name: "cannot baseline a missing sql file",
			args: args{
				path:     absolutePath,
				fileName: "missing.sql",
				fileType: params.FileTypes["-f"],
			},
			expectedMessage: "",
			expectedErr:     FileNotFoundError,
		},
		{
			name: "cannot baseline a non existent source path",
			args: args{
				path:     filepath.Join(absolutePath, "nonexistent"),
				fileName: "missing.sql",
				fileType: params.FileTypes["-f"],
			},
			expectedMessage: "",
			expectedErr:     SourcePathError,
		},
	}
	for _, tt := range tests {
		got, err := Baseline(tt.args.path, tt.args.fileName, tt.args.fileType)

		assert.Equal(t, tt.expectedErr, errors.Context(err))
		assert.Equal(t, tt.expectedMessage, got)

		t.Cleanup(func() {
			_ = os.RemoveAll(filepath.Join(absolutePath, "release"))
		})
	}
}
