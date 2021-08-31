package main

import (
	"context"

	"github.com/alecthomas/kong"
	"github.com/danielMensah/sqlr/internal/command"
)

var cli struct {
	Baseline command.Baseline `cmd:"" name:"baseline" help:"Baseline new sql file"`
	Deploy   command.Deploy   `cmd:"" name:"deploy" help:"Deploy sql files"`
}

func main() {
	ctx := kong.Parse(
		&cli,
		kong.Description("A collection of commands for SQL deployments with version control"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)

	runCtx := &command.Context{
		Context: context.Background(),
	}

	err := ctx.Run(runCtx)
	ctx.FatalIfErrorf(err)
}

//func main() {
//	path, err := os.Getwd()
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	switch option := os.Args[1]; option {
//	case "baseline":
//		//b := baseline.New(path, os.Args[2], os.Args[3])
//		//
//		//e := b.BaselineFile()
//		//if e != nil {
//		//	logrus.Error(e)
//		//}
//		//fileTypeParam := os.Args[2]
//		//fileName := os.Args[3]
//		//
//		//if fileTypeParam == "" {
//		//	logrus.Error(errors.New("missing file type parameter"))
//		//	return
//		//} else if fileName == "" {
//		//	logrus.Error(errors.New("missing file name parameter"))
//		//	return
//		//}
//		//
//		//fileTypeParam = params.FileTypes[fileTypeParam]
//		//
//		//response, responseErr := directory.Baseline(path, fileName, fileTypeParam)
//		//
//		//if responseErr != nil {
//		//	logrus.Error(responseErr)
//		//}
//		//
//		//logrus.Println(response)
//	case "new":
//		currentPath := utility.GetCurrentPath(path, []string{})
//		err = directory.CreateFolder(currentPath)
//
//		if err != nil {
//			logrus.Error(err)
//		}
//
//		logrus.Info("Current created!")
//	case "help":
//		utility.ShowHelp()
//	case "release":
//
//	default:
//		fmt.Println("option not found")
//	}
//}
