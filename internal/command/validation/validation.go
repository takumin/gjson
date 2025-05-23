package validation

import (
	"context"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gjson/internal/config"
	"github.com/takumin/gjson/internal/filelist"
	"github.com/takumin/gjson/internal/parser"
	"github.com/takumin/gjson/internal/report"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.StringSliceFlag{
			Name:        "include",
			Aliases:     []string{"i"},
			Usage:       "include file extension",
			Sources:     cli.EnvVars("INCLUDE"),
			Value:       cfg.Extention.Includes,
			Destination: &cfg.Extention.Includes,
		},
		&cli.StringSliceFlag{
			Name:        "exclude",
			Aliases:     []string{"e"},
			Usage:       "exclude file extension",
			Sources:     cli.EnvVars("EXCLUDE"),
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

func before(cfg *config.Config) func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
	return func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
		if cmd.NArg() > 0 {
			s := cmd.Args().Slice()
			sort.Strings(s)
			cfg.Path.Searches = slices.Compact(s)
		}
		return ctx, nil
	}
}

func action(cfg *config.Config) func(ctx context.Context, cmd *cli.Command) error {
	return func(ctx context.Context, cmd *cli.Command) error {
		paths := make([]string, 0, 65535)
		for _, path := range cfg.Path.Searches {
			info, err := os.Stat(path)
			if err != nil {
				return err
			}

			if info.IsDir() {
				list, err := filelist.Filelist(
					os.DirFS(path),
					path,
					cfg.Extention.Includes,
					cfg.Extention.Excludes,
				)
				if err != nil {
					return err
				}
				if len(list) > 0 {
					paths = append(paths, list...)
				}
			} else {
				paths = append(paths, path)
			}
		}

		perrs := make([]*parser.ParseError, 0, len(paths))
		for _, path := range paths {
			perr, err := parser.Parse(path)
			if err != nil {
				return err
			}
			if perr != nil {
				perrs = append(perrs, perr)
			}
		}

		var buf strings.Builder
		for _, perr := range perrs {
			res, err := report.ReviewdogDiagnosticJSONLines(perr.Filename, perr.Message, perr.Line, perr.Column)
			if err != nil {
				return err
			}
			buf.Write(res)
			buf.WriteString("\n")
		}

		if buf.Len() > 0 {
			fmt.Print(buf.String())
			os.Exit(2)
		}

		return nil
	}
}
