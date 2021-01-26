package main

import (
	"fmt"
	"github.com/danielMensah/sqlr/internal/directory"
	"github.com/danielMensah/sqlr/internal/params"
	"github.com/danielMensah/sqlr/internal/utility"
	"github.com/danielMensah/sqlr/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	path, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	switch option := os.Args[1]; option {
	case "baseline":
		fileTypeParam := os.Args[2]
		fileName := os.Args[3]

		if fileTypeParam == "" {
			logrus.Error(errors.New("missing file type parameter"))
			return
		} else if fileName == "" {
			logrus.Error(errors.New("missing file name parameter"))
			return
		}

		fileTypeParam = params.FileTypes[fileTypeParam]

		response, responseErr := directory.Baseline(path, fileName, fileTypeParam)

		if responseErr != nil {
			logrus.Error(responseErr)
		}

		logrus.Println(response)
	case "new":
		currentPath := utility.GetCurrentPath(path, []string{})
		err = directory.CreateFolder(currentPath)

		if err != nil {
			logrus.Error(err)
		}

		logrus.Info("Current created!")
	default:
		fmt.Println("option not found")
	}
}
