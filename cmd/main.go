package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewGoApp() *cli.App {
	return &cli.App{
		Name:  "ngrok-go-example",
		Usage: "A simple service written to use the Ngrok Go networking ingress",
		Commands: []*cli.Command{
			NewServerCommand(),
		},
	}
}
