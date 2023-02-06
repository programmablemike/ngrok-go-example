package main

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/programmablemike/ngrok-go-example/cmd"
)

func main() {
	app := cmd.NewGoApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err)
	}
}
