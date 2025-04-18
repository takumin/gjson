package fish

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gjson/internal/config"
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
