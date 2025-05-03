package powershell

import (
	"context"
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gjson/internal/config"
)

const powershellCompletion = `
$fn = $($MyInvocation.MyCommand.Name)
$name = $fn -replace "(.*)\.ps1$", '$1'
Register-ArgumentCompleter -Native -CommandName $name -ScriptBlock {
	param($commandName, $wordToComplete, $cursorPosition)
	$other = "$wordToComplete --generate-shell-completion"
	Invoke-Expression $other | ForEach-Object {
		[System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
	}
}
`

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "powershell",
		Usage:    "powershell completion",
		HideHelp: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println(strings.TrimSpace(powershellCompletion))
			return nil
		},
	}
}
