package directory

import (
	"fmt"
	"io"
	"os"
)

// CreateFolder creates sub-directories if they do not exist
func CreateFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0700)

		if err != nil {
			return fmt.Errorf("creating folders %s: %w", path, err)
		}
	}

	return nil
}

// CopyToDestination copies files from source to destination
func CopyToDestination(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("get source file content %s: %w", source, err)
	}
	defer sourceFile.Close()

	newFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("file creation %s: %w", destination, err)
	}
	defer newFile.Close()

	_, copyError := io.Copy(newFile, sourceFile)
	if copyError != nil {
		return fmt.Errorf("copy file %s: %w", source, err)
	}

	return nil
}
