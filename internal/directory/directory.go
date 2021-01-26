package directory

import (
	"github.com/danielMensah/sqlr/internal/utility"
	"github.com/danielMensah/sqlr/pkg/errors"
	"io"
	"os"
	"path/filepath"
)

const BaselineResponse = "Baseline done!"

var (
	FileNotFoundError    = errors.New("file not found in source folder")
	NewFileCreationError = errors.New("new file could not be created in target folder")
	FolderCreationError  = errors.New("failed creating folder")
	FileCopyError        = errors.New("file could not be copied in target folder")
	SourcePathError      = errors.New("source folder not found")
)

func Baseline(path string, fileName string, fileType string) (string, error) {
	// No need to create folders. Schema should already be created
	source := filepath.Join(path, "schema", fileType)

	if _, err := os.Stat(source); os.IsNotExist(err) {
		return "", SourcePathError.WithCause(err)
	}

	target := utility.GetCurrentPath(path, []string{fileType})
	err := CreateFolder(target)

	if err != nil {
		return "", FolderCreationError.WithCause(err)
	}

	err = copyToCurrent(filepath.Join(source, fileName), filepath.Join(target, fileName))

	if err != nil {
		return "", err
	}

	return BaselineResponse, nil
}

func CreateFolder(path string) error {
	// Create sub-directories if they do not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0700)

		if err != nil {
			return FolderCreationError.WithCause(err)
		}
	}

	return nil
}

func copyToCurrent(source string, target string) error {
	// Get source file
	sourceFile, err := os.Open(source)
	if err != nil {
		return FileNotFoundError.WithCause(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, newFileError := os.Create(target)
	if newFileError != nil {
		return NewFileCreationError.WithCause(newFileError)
	}
	defer newFile.Close()

	// Copy source file content to target file
	_, copyError := io.Copy(newFile, sourceFile)
	if copyError != nil {
		return FileCopyError.WithCause(copyError)
	}

	return nil
}
