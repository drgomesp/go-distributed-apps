package main

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/drgomesp/go-distributed-apps/02-configuration/internal/config"
)

var cfg config.Config

func init() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var err error
	cfg, err = env.ParseAs[config.Config]()
	if err != nil {
		log.Err(err).Msg("could not parse environment variables")
	}
}

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make  an explosive entrance",
		Action: func(c *cli.Context) error {
			log.Info().Msg("hello world!")

			spew.Dump(cfg)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err)
	}
}