package fish

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/takumin/boilerplate-golang-cli/internal/config"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "fish",
		Usage:    "fish completion",
		HideHelp: true,
		Action: func(ctx *cli.Context) error {
			fish, err := ctx.App.ToFishCompletion()
			if err != nil {
				return err
			}
			fmt.Println(fish)
			return nil
		},
	}
}
