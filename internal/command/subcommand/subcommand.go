package subcommand

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/boilerplate-golang-cli/internal/config"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "variable",
			Aliases:     []string{"v"},
			Usage:       "variable",
			EnvVars:     []string{"VARIABLE"},
			Value:       cfg.Variable,
			Destination: &cfg.Variable,
		},
	}...)
	return &cli.Command{
		Name:    "subcommand",
		Aliases: []string{"sub"},
		Usage:   "subcommand",
		Flags:   flags,
		Action:  action(cfg),
	}
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return nil
	}
}
