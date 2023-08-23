package validation

import (
	"fmt"
	"os"
	"slices"
	"sort"

	"github.com/urfave/cli/v2"

	"github.com/takumin/gjson/internal/config"
	"github.com/takumin/gjson/internal/filelist"
	"github.com/takumin/gjson/internal/validate"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.MultiStringFlag{
			Target: &cli.StringSliceFlag{
				Name:    "include",
				Aliases: []string{"i"},
				Usage:   "include file extension",
				EnvVars: []string{"INCLUDE"},
			},
			Value:       cfg.Extention.Includes,
			Destination: &cfg.Extention.Includes,
		},
		&cli.MultiStringFlag{
			Target: &cli.StringSliceFlag{
				Name:    "exclude",
				Aliases: []string{"e"},
				Usage:   "exclude file extension",
				EnvVars: []string{"EXCLUDE"},
			},
			Value:       cfg.Extention.Excludes,
			Destination: &cfg.Extention.Excludes,
		},
	}...)
	return &cli.Command{
		Name:            "validation",
		Aliases:         []string{"validate", "valid", "v"},
		Usage:           "json file validation",
		ArgsUsage:       "[file or directory...]",
		HideHelpCommand: true,
		Flags:           flags,
		Before:          before(cfg),
		Action:          action(cfg),
	}
}

func before(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		if ctx.NArg() > 0 {
			s := ctx.Args().Slice()
			sort.Strings(s)
			cfg.Path.Searches = slices.Compact(s)
		}
		return nil
	}
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		paths := make([]string, 0, 65535)

		for _, path := range cfg.Path.Searches {
			list, err := filelist.Filelist(
				os.DirFS(path),
				cfg.Path.RootDir,
				cfg.Extention.Includes,
				cfg.Extention.Excludes,
			)
			if err != nil {
				return err
			}
			if len(list) > 0 {
				paths = append(paths, list...)
			}
		}

		// TODO: refactoring
		if res, err := validate.Validate(paths); err != nil {
			for _, v := range res {
				fmt.Fprintln(os.Stderr, v)
			}
			os.Exit(2)
		}

		return nil
	}
}
