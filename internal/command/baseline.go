package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/danielMensah/sqlr/internal/directory"
)

const (
	SourcePath      = "schema"
	DestinationPath = "current/schema"
)

type Baseline struct {
	Folder      string `arg:"" required:"" help:"Folder where the sql file resides"`
	FileName    string `arg:"" required:"" help:"SQL file name to baseline"`
	Source      string
	Destination string
}

func (b Baseline) Run(ctx *Context) error {
	err := b.Setup()
	if err != nil {
		return err
	}

	return directory.CopyToDestination(b.Source, b.Destination)
}

func (b *Baseline) Setup() error {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	source := filepath.Join(path, SourcePath, b.Folder)
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return fmt.Errorf("initialise source: %w", err)
	}

	b.Source = source
	b.Destination = filepath.Join(path, DestinationPath, b.Folder)

	return directory.CreateFolder(b.Destination)
}
