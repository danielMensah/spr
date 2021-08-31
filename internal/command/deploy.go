package command

type Deploy struct {
	Stage string `required short:"s" help:"Deployment environment (dev, demo, prod)"`
}

func (d Deploy) Run() error {

	return nil
}
