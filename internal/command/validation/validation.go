package validation

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/takumin/gjson/internal/config"
	"github.com/takumin/gjson/internal/filelist"
	"github.com/takumin/gjson/internal/validate"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "directory",
			Aliases:     []string{"d"},
			Usage:       "search base directory",
			EnvVars:     []string{"DIRECTORY"},
			Value:       cfg.Path.Directory,
			Destination: &cfg.Path.Directory,
		},
	}...)
	return &cli.Command{
		Name:    "validation",
		Aliases: []string{"validate", "valid", "v"},
		Usage:   "validation json files",
		Flags:   flags,
		Action:  action(cfg),
	}
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		if ctx.Args().Len() == 1 {
			cfg.Path.Directory = ctx.Args().First()
		}

		list, err := filelist.Filelist(
			cfg.Path.Directory,
			strings.Split(cfg.Extention.Includes, ","),
			strings.Split(cfg.Extention.Excludes, ","),
		)
		if err != nil {
			return err
		}

		// TODO: refactoring
		if res, err := validate.Validate(list); err != nil {
			for _, v := range res {
				fmt.Fprintln(os.Stderr, v)
			}
			os.Exit(2)
		}

		return nil
	}
}
