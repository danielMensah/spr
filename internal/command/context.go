package command

import (
	"context"
)

// GlobalArgs is the available flags present in every command
type GlobalArgs struct {
	Stage string `required short:"s" help:"Deployment environment (dev, demo, prod)"`
}

// Context carries values to be passed to every executed command
type Context struct {
	GlobalArgs
	context.Context
}
