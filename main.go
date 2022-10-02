package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/takumin/boilerplate-golang-cli/internal/command/completion"
	"github.com/takumin/boilerplate-golang-cli/internal/command/subcommand"
	"github.com/takumin/boilerplate-golang-cli/internal/config"
)

var (
	AppName  string = "boilerplate-golang-cli"      // ###BOILERPLATE_APP_NAME###
	AppDesc  string = "Boilerplate Golang CLI Tool" // ###BOILERPLATE_APP_DESC###
	Version  string = "unknown"
	Revision string = "unknown"
)

func main() {
	cfg := config.NewConfig()

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "log-level",
			Aliases:     []string{"l"},
			Usage:       "log level",
			EnvVars:     []string{"LOG_LEVEL"},
			Value:       cfg.LogLevel,
			Destination: &cfg.LogLevel,
		},
	}

	cmds := []*cli.Command{
		completion.NewCommands(cfg, flags),
		subcommand.NewCommands(cfg, flags),
	}

	app := &cli.App{
		Name:                 AppName,
		Usage:                AppDesc,
		Version:              fmt.Sprintf("%s (%s)", Version, Revision),
		Flags:                flags,
		Commands:             cmds,
		EnableBashCompletion: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
